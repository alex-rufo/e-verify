package sanetization

import (
	"testing"
)

var test = []struct {
	name         string
	emailAddress string
	expected     string
}{
	{"not gmail", "alex.rufo+3@hotmail.com", "alex.rufo+3@hotmail.com"},
	{"gmail with +", "alexrufo+2@gmail.com", "alexrufo@gmail.com"},
	{"gmail with .", "alex.rufo@gmail.com", "alexrufo@gmail.com"},
	{"gmail with . and +", "alex.rufo+hello@gmail.com", "alexrufo@gmail.com"},
}

func TestItSanitizesValidEmails(t *testing.T) {
	s := &Gmail{}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			sanetizedEmail, _ := s.Sanitize(test.emailAddress)
			if sanetizedEmail != test.expected {
				t.Errorf("Invalid gmail sanity: %s. Expected: %s", sanetizedEmail, test.expected)
			}
		})
	}
}

func TestItReturnsAnErrorIfItsAnInvalidEmail(t *testing.T) {
	s := &Gmail{}
	_, err := s.Sanitize("this is not an email")
	if err == nil {
		t.Error("An error should be returned if the email is invalid")
	}
}
