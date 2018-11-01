package sanetization

import "strings"

// Lowercase sanetization
type Lowercase struct{}

// Sanitize converts the email to lowercase
func (l *Lowercase) Sanitize(email string) (string, error) {
	return strings.ToLower(email), nil
}
