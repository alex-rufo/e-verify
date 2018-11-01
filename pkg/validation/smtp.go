package validation

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/alex-rufo/e-verify/internal/email"
)

// Dialer is the interface for our own Dialer
type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type SMTP struct {
	dialer Dialer
}

func NewSMTP(d Dialer) *SMTP {
	return &SMTP{dialer: d}
}

func (s *SMTP) Validate(emailAddress string) (bool, error) {
	domain, err := email.GetDomain(emailAddress)
	if err != nil {
		return false, err
	}

	conn, err := s.dialer.Dial("tcp", fmt.Sprintf("%s:%d", domain, 25))
	if err != nil {
		return false, err
	}
	defer conn.Close()
	defer fmt.Fprintf(conn, "quit\n")

	bufio.NewReader(conn).ReadString('\n')
	fmt.Fprintf(conn, fmt.Sprintf("helo %s\n", domain))
	bufio.NewReader(conn).ReadString('\n')
	fmt.Fprintf(conn, fmt.Sprintf("mail from: <%s>\n", emailAddress))
	bufio.NewReader(conn).ReadString('\n')
	fmt.Fprintf(conn, fmt.Sprintf("rcpt to: <%s>\n", emailAddress))
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return false, err
	}

	return !strings.Contains(status, "250"), nil
}
