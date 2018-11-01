package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	verifier "github.com/alex-rufo/e-verify"
	"github.com/alex-rufo/e-verify/proto"
	"github.com/spf13/cobra"
	ctx "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var grpcPort int

func init() {
	rootCmd.PersistentFlags().IntVarP(&grpcPort, "grpc-port", "", 81, "Port for gRPC server")
	rootCmd.AddCommand(rpcCmd)
}

var rpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Starts a gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
		if err != nil {
			log.Fatal(err)
		}

		grpcServer := grpc.NewServer()
		adapter := &VerifyAdapter{verifier: emailVerifier}
		proto.RegisterVerifyServiceServer(grpcServer, adapter)

		go func() {
			if err := grpcServer.Serve(listener); err != nil {
				log.Fatal(err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		grpcServer.GracefulStop()
	},
}

// VerifyAdapter adapts the gRPC interface into the verify service one
type VerifyAdapter struct {
	verifier *verifier.Verifier
}

// Verify adapts the request and response objects
func (a *VerifyAdapter) Verify(ctx ctx.Context, request *proto.EmailVerifyRequest) (*proto.EmailVerifyResponse, error) {
	valid, err := a.verifier.Verify(request.GetEmail())
	return &proto.EmailVerifyResponse{Valid: valid, Error: err.Error()}, nil
}
