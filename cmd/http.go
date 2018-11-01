package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

type Response struct {
	Valid bool `json:"valid"`
}

var httpPort int

func init() {
	rootCmd.PersistentFlags().IntVarP(&httpPort, "http-port", "", 80, "Port for HTTP server")
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts an HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()
		e.GET("/verify/:email", func(c echo.Context) error {
			valid, err := emailVerifier.Verify(c.Param("email"))
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, err)
			}
			return c.JSON(http.StatusOK, Response{Valid: valid})
		})

		go func() {
			if err := e.Start(fmt.Sprintf(":%d", httpPort)); err != nil {
				log.Fatal(err)
			}
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	},
}
