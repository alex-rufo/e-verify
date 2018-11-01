package validation

import (
	"testing"
)

func TestRoleValidation(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"alfons@domain.com", true},
		{"alex@gmail.com", true},
		{"jordi@hotmail.com", true},
		{"2018@yahoo.com", false},
		{"admin@mailinator.com", false},
		{"advertisingsales@guerrillamail.com", false},
		{"accounts@yopmail.com", false},
	}

	v := NewRole([]string{"2018", "admin", "advertisingsales", "accounts"})
	for _, test := range tests {
		valid, _ := v.Validate(test.email)
		if valid != test.expected {
			t.Errorf("The expected validation for %s was %t but the got %t", test.email, test.expected, valid)
		}
	}
}

func TestRoleValidationReturnsAnErrorIfEmailIsInvalid(t *testing.T) {
	v := NewRole([]string{})
	_, err := v.Validate("invalid email")
	if err == nil {
		t.Error("An error should be returned if the email is invalid")
	}
}
