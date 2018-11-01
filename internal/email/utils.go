package email

import (
	"errors"
	"strings"
)

// ErrInvalidEmail is thrown if the email is not valid
var ErrInvalidEmail = errors.New("Invalid email address")

// GetLocalAndDomain from a given email address
func GetLocalAndDomain(email string) (string, string, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", "", ErrInvalidEmail
	}

	return parts[0], parts[1], nil
}

// GetLocal from a given email address
func GetLocal(email string) (string, error) {
	local, _, err := GetLocalAndDomain(email)
	return local, err
}

// GetDomain from a given email address
func GetDomain(email string) (string, error) {
	_, domain, err := GetLocalAndDomain(email)
	return domain, err
}
