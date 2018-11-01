package validation

import (
	"github.com/alex-rufo/e-verify/internal/email"
	"github.com/alex-rufo/e-verify/internal/parse"
)

type Domain struct {
	domains []string
}

func NewDomain(domains []string) *Domain {
	return &Domain{domains: domains}
}

func NewDomainFromFile(file string) (*Domain, error) {
	domains, err := parse.From(file)
	if err != nil {
		return nil, err
	}

	return NewDomain(domains), nil
}

// Validate if the email is domain is disposable
func (d *Domain) Validate(emailAddress string) (bool, error) {
	domain, err := email.GetDomain(emailAddress)
	if err != nil {
		return false, err
	}

	for _, d := range d.domains {
		if domain == d {
			return false, nil
		}
	}

	return true, nil
}
