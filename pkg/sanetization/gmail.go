package sanetization

import (
	"fmt"
	"strings"

	"github.com/alex-rufo/e-verify/internal/email"
)

// Gmail sanitizes specific gmail emails
type Gmail struct{}

// Sanitize matches the interface
func (g *Gmail) Sanitize(emailAddress string) (string, error) {
	local, domain, err := email.GetLocalAndDomain(emailAddress)
	if err != nil {
		return "", err
	}

	if domain != "gmail.com" {
		return emailAddress, nil
	}

	local = strings.Replace(local, ".", "", -1)
	index := strings.Index(local, "+")
	if index > -1 {
		local = local[0:index]
	}
	return fmt.Sprintf("%s@%s", local, domain), nil
}
