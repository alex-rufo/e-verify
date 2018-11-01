package validation

import (
	"context"
	"net"
	"time"

	"github.com/alex-rufo/e-verify/internal/email"
)

type MXResolver interface {
	LookupMX(ctx context.Context, domain string) ([]*net.MX, error)
}

type MX struct {
	resolver MXResolver
	timeout  time.Duration
}

func NewMX(r MXResolver, t time.Duration) *MX {
	return &MX{
		resolver: r,
		timeout:  t,
	}
}

func (m *MX) Validate(emailAddress string) (bool, error) {
	domain, err := email.GetDomain(emailAddress)
	if err != nil {
		return false, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	mx, err := m.resolver.LookupMX(ctx, domain)
	if err != nil {
		return false, nil
	}

	return len(mx) != 0, nil
}
