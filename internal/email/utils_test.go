package email

import "testing"

func TestItGetsTheDomain(t *testing.T) {
	domain, err := GetDomain("local@domain.com")
	if err != nil {
		t.Error("No error should be returned if the email is valid")
	}
	if domain != "domain.com" {
		t.Errorf("The domain of 'local@domain.com' should be 'domain.com', given %s", domain)
	}
}

func TestItGetsTheLocal(t *testing.T) {
	local, err := GetLocal("local@domain.com")
	if err != nil {
		t.Error("No error should be returned if the email is valid")
	}
	if local != "local" {
		t.Errorf("The local of 'local@domain.com' should be 'local', given %s", local)
	}
}

func TestItGetsTheDomainAndLocal(t *testing.T) {
	local, domain, err := GetLocalAndDomain("local@domain.com")
	if err != nil {
		t.Error("No error should be returned if the email is valid")
	}
	if local != "local" {
		t.Errorf("The local of 'local@domain.com' should be 'local', given %s", local)
	}
	if domain != "domain.com" {
		t.Errorf("The domain of 'local@domain.com' should be 'domain.com', given %s", domain)
	}
}

func TestItReturnsAnErrorIfTheEmailIsNotValid(t *testing.T) {
	_, _, err := GetLocalAndDomain("invalid domain")
	if err == nil {
		t.Error("An error should be returned if the email is not valid")
	}
}
