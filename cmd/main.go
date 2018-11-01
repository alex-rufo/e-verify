package main

import (
	"log"
	"net"
	"time"

	verifier "github.com/alex-rufo/e-verify"
	"github.com/alex-rufo/e-verify/pkg/sanetization"
	"github.com/alex-rufo/e-verify/pkg/validation"
	"github.com/spf13/cobra"
)

var (
	disposableDomainsFile string
	disposableRolesFile   string
	xverifyAPIKey         string
	xverifyDomain         string
	enableSMTP            bool

	emailVerifier *verifier.Verifier
)

var rootCmd = &cobra.Command{
	Use:     "e-verify",
	Short:   "e-verify cli",
	Version: "0.0.1",
}

// Execute the command from the CLI
func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initialize)
	rootCmd.PersistentFlags().StringVarP(&disposableDomainsFile, "disposable-domains", "", "", "File where disposable domains are located")
	rootCmd.PersistentFlags().StringVarP(&disposableRolesFile, "disposable-roles", "", "", "File where disposable roles are located")
	rootCmd.PersistentFlags().StringVarP(&xverifyAPIKey, "xverify-apiKey", "", "", "ApiKey of xVerify")
	rootCmd.PersistentFlags().StringVarP(&xverifyDomain, "xverify-domain", "", "", "Domain to use with xVerify")
	rootCmd.PersistentFlags().BoolVarP(&enableSMTP, "enable-smtp", "", false, "Enable SMTP?")
}

func initialize() {
	sanitizers := []verifier.Sanitizer{
		&sanetization.Trim{},
		&sanetization.Lowercase{},
		&sanetization.Gmail{},
	}

	validators := []verifier.Validator{
		&validation.Syntax{},
		validation.NewMX(&net.Resolver{}, 1*time.Second),
	}
	if disposableDomainsFile != "" {
		validator, err := validation.NewDomainFromFile(disposableDomainsFile)
		if err != nil {
			log.Fatal(err)
		}
		validators = append(validators, validator)
	}
	if disposableRolesFile != "" {
		validator, err := validation.NewRoleFromFile(disposableRolesFile)
		if err != nil {
			log.Fatal(err)
		}
		validators = append(validators, validator)
	}
	if enableSMTP {
		validators = append(validators, validation.NewSMTP(&net.Dialer{Timeout: 1 * time.Second}))
	}
	if xverifyAPIKey != "" && xverifyDomain != "" {
		validators = append(validators, validation.NewXVerify(xverifyAPIKey, xverifyDomain, 1*time.Second))
	}

	emailVerifier = verifier.New(sanitizers, validators)
}
