package everify

import (
	"net"
	"time"

	"github.com/alex-rufo/e-verify/pkg/sanetization"
	"github.com/alex-rufo/e-verify/pkg/validation"
)

// Sanitizer interface
type Sanitizer interface {
	Sanitize(email string) (string, error)
}

// Validator interface
type Validator interface {
	Validate(email string) (bool, error)
}

// Verifier implementation
type Verifier struct {
	sanitizers []Sanitizer
	validators []Validator
}

// New returns a new verifier instance
func New(s []Sanitizer, v []Validator) *Verifier {
	return &Verifier{
		sanitizers: s,
		validators: v,
	}
}

// NewDefault returns a new verifier with the default sanitizer and validator
func NewDefault() *Verifier {
	sanitizers := []Sanitizer{
		&sanetization.Trim{},
		&sanetization.Lowercase{},
		&sanetization.Gmail{},
	}

	domainValidator, _ := validation.NewDomainFromFile("https://raw.githubusercontent.com/martenson/disposable-email-domains/master/allowlist.conf")
	validators := []Validator{
		&validation.Syntax{},
		domainValidator,
		validation.NewMX(&net.Resolver{}, 1*time.Second),
	}

	return New(sanitizers, validators)
}

// Verify validates if the email address is valid
func (v *Verifier) Verify(email string) (bool, error) {
	sanitizedEmail, err := v.sanitize(email)
	if err != nil {
		return false, err
	}

	return v.validate(sanitizedEmail)
}

// Sanitizes the email
func (v *Verifier) sanitize(email string) (string, error) {
	var err error
	for _, sanitizer := range v.sanitizers {
		email, err = sanitizer.Sanitize(email)
		if err != nil {
			return email, err
		}
	}

	return email, nil
}

// Validates the email
func (v *Verifier) validate(email string) (bool, error) {
	for _, validator := range v.validators {
		valid, err := validator.Validate(email)
		if err != nil {
			return false, err
		}
		if !valid {
			return false, nil
		}
	}

	return true, nil
}
