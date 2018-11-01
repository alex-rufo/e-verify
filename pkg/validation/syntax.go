package validation

import (
	"regexp"
)

const emailSyntaxRegexp = "\\A[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\z"

type Syntax struct{}

// Validate if the email has a valid syntax
func (s *Syntax) Validate(email string) (bool, error) {
	regexp := regexp.MustCompile(emailSyntaxRegexp)
	return regexp.MatchString(email), nil
}
