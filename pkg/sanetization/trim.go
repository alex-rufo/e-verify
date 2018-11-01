package sanetization

import "strings"

// Trim sanetization
type Trim struct{}

// Sanitize converts the email to lowercase
func (t *Trim) Sanitize(email string) (string, error) {
	return strings.TrimSpace(email), nil
}
