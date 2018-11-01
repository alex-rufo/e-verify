package everify

import (
	"errors"
	"testing"
)

type MockSanitizer struct{}

func (m *MockSanitizer) Sanitize(email string) (string, error) {
	return "", errors.New("test error")
}

type MockValidator struct {
	called     bool
	valideFunc func(email string) (bool, error)
}

func (m *MockValidator) Validate(email string) (bool, error) {
	m.called = true
	if m.valideFunc != nil {
		return m.valideFunc(email)
	}
	return true, nil
}
func (m *MockValidator) HasBeenCalled() bool {
	return m.called
}

func TestItReturnsAnErrorIfCanNotSanitize(t *testing.T) {
	v := New([]Sanitizer{&MockSanitizer{}}, []Validator{})
	_, err := v.Verify("email@email.com")
	if err == nil {
		t.Error("An error should be returned if fails sanitizing")
	}
}

func TestItReturnsAnErrorIfCanValidate(t *testing.T) {
	validator := &MockValidator{
		valideFunc: func(email string) (bool, error) {
			return false, errors.New("test error")
		},
	}

	v := New([]Sanitizer{}, []Validator{validator})
	_, err := v.Verify("email@email.com")
	if err == nil {
		t.Error("An error should be returned if fails validating")
	}
}

func TestItStopsValidatingIfItIsNotValid(t *testing.T) {
	validator1 := &MockValidator{
		valideFunc: func(email string) (bool, error) {
			return false, nil
		},
	}
	validator2 := &MockValidator{}

	v := New([]Sanitizer{}, []Validator{validator1, validator2})
	valid, err := v.Verify("email@email.com")
	if err != nil {
		t.Error("No error should be returned")
	}

	if valid {
		t.Error("The response should be false if one of the validator says it's invalid")
	}
	if validator2.HasBeenCalled() {
		t.Error("The chain shoudl be stopped if the email is invalid")
	}
}

func TestItCompletesTheChainIfIsValid(t *testing.T) {
	validator1 := &MockValidator{}
	validator2 := &MockValidator{}

	v := New([]Sanitizer{}, []Validator{validator1, validator2})
	valid, err := v.Verify("email@email.com")
	if err != nil {
		t.Error("No error should be returned")
	}

	if !valid {
		t.Error("The response should be true if all the validators say so")
	}
	if !validator1.HasBeenCalled() || !validator2.HasBeenCalled() {
		t.Error("All the validator should have been called if the email is valid")
	}
}
