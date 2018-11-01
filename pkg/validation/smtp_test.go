package validation

import (
	"errors"
	"net"
	"testing"
)

type MockDialer struct {
	dialFunc func(network, address string) (net.Conn, error)
}

func (m *MockDialer) Dial(network, address string) (net.Conn, error) {
	if m.dialFunc != nil {
		return m.dialFunc(network, address)
	}

	return nil, nil
}

func TestSMTPValidationReturnsErrorIfTheEmailidInvalid(t *testing.T) {
	v := NewSMTP(&MockDialer{})
	_, err := v.Validate("invalid email")
	if err == nil {
		t.Error("An error should be returned if the email is invalid")
	}
}

func TestSMTPValidationReturnsErrorIfCanNotDial(t *testing.T) {
	dialer := &MockDialer{
		dialFunc: func(network, address string) (net.Conn, error) {
			return nil, errors.New("test error")
		},
	}

	v := NewSMTP(dialer)
	_, err := v.Validate("email@email.com")
	if err == nil {
		t.Error("An error should be returned if can not dial")
	}
}
