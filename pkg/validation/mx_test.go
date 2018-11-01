package validation

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"
)

type MockMXResolver struct {
	lookupMXFunc func(ctx context.Context, domain string) ([]*net.MX, error)
}

func (m *MockMXResolver) LookupMX(ctx context.Context, domain string) ([]*net.MX, error) {
	if m.lookupMXFunc != nil {
		return m.lookupMXFunc(ctx, domain)
	}

	return []*net.MX{}, nil
}

func TestMXValidationReturnsAnErrorIfEmailIsInvalid(t *testing.T) {
	v := NewMX(&MockMXResolver{}, time.Second)
	_, err := v.Validate("invalid email")
	if err == nil {
		t.Error("An error should be returned if the email is invalid")
	}
}

func TestMXValidationReturnsInvalidIfLookupMXFails(t *testing.T) {
	resolver := &MockMXResolver{
		lookupMXFunc: func(ctx context.Context, domain string) ([]*net.MX, error) {
			return []*net.MX{}, errors.New("test error")
		},
	}

	v := NewMX(resolver, time.Second)
	valid, err := v.Validate("email@email.com")
	if err != nil {
		t.Error("No error should be returned if lookup does not find the host")
	}
	if valid {
		t.Error("The email should no be valid if lookup returns an error")
	}
}

func TestMXValidationReturnsFalseIfNoMXIsFound(t *testing.T) {
	resolver := &MockMXResolver{
		lookupMXFunc: func(ctx context.Context, domain string) ([]*net.MX, error) {
			return []*net.MX{}, nil
		},
	}

	v := NewMX(resolver, time.Second)
	valid, err := v.Validate("email@email.com")
	if err != nil {
		t.Error("No error should be thrown")
	}
	if valid {
		t.Error("The email should no be valid if no MX is found")
	}
}

func TestMXValidationReturnsTrueIfAtLeastOneMXIsFound(t *testing.T) {
	resolver := &MockMXResolver{
		lookupMXFunc: func(ctx context.Context, domain string) ([]*net.MX, error) {
			return []*net.MX{&net.MX{}}, nil
		},
	}

	v := NewMX(resolver, time.Second)
	valid, err := v.Validate("email@email.com")
	if err != nil {
		t.Error("No error should be thrown")
	}
	if !valid {
		t.Error("The email should be valid if at least one MX is found")
	}
}
