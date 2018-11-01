package validation

import (
	"testing"
)

func TestSyntaxValidation(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"email@domain.com", true},
		{"firstname.lastname@domain.com", true},
		{"email@subdomain.domain.com", true},
		{"firstname+lastname@domain.com", true},
		{"email@123.123.123.123", true},
		{"1234567890@domain.com", true},
		{"email@domain-one.com", true},
		{"_______@domain.co", true},
		{"email@domain.name", true},
		{"email@domain.co.jp", true},
		{"firstname-lastname@domain.com", true},
		{"plainaddress", false},
		{"#@%^%#$@#$@#.com", false},
		{"@domain.com", false},
		{"Joe Smith <email@domain.com>", false},
		{"email.domain.com", false},
		{"email@domain@domain.com", false},
		{".email@domain.com", false},
		{"email.@domain.com", false},
		{"email..email@domain.com", false},
		{"あいうえお@domain.com", false},
		{"email@domain.com (Joe Smith)", false},
		{"\"email\"@domain.com", false},
		{"email@[123.123.123.123]", false},
		{"email@domain", false},
		{"email@-domain.com", false},
		{"email@domain..com", false},
	}

	s := &Syntax{}
	for _, test := range tests {
		valid, _ := s.Validate(test.email)
		if valid != test.expected {
			t.Errorf("The expected validation for %s was %t but the got %t", test.email, test.expected, valid)
		}
	}
}
