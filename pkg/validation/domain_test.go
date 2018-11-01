package validation

import (
	"testing"
)

func TestDisposableValidation(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"email@domain.com", true},
		{"email@gmail.com", true},
		{"email@hotmail.com", true},
		{"email@yahoo.com", true},
		{"email@mailinator.com", false},
		{"email@guerrillamail.com", false},
	}

	v := NewDomain([]string{"mailinator.com", "guerrillamail.com"})
	for _, test := range tests {
		valid, _ := v.Validate(test.email)
		if valid != test.expected {
			t.Errorf("The expected validation for %s was %t but the got %t", test.email, test.expected, valid)
		}
	}
}

func TestDisposableValidationReturnsAnErrorIfEmailIsInvalid(t *testing.T) {
	v := NewDomain([]string{})
	_, err := v.Validate("invalid email")
	if err == nil {
		t.Error("An error should be returned if the email is invalid")
	}
}
