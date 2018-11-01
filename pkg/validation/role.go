package validation

import (
	"github.com/alex-rufo/e-verify/internal/email"
	"github.com/alex-rufo/e-verify/internal/parse"
)

type Role struct {
	roles []string
}

func NewRole(roles []string) *Role {
	return &Role{roles: roles}
}

func NewRoleFromFile(file string) (*Role, error) {
	roles, err := parse.From(file)
	if err != nil {
		return nil, err
	}

	return NewRole(roles), nil
}

// Validate if the email address local is a role
func (r *Role) Validate(emailAddress string) (bool, error) {
	local, err := email.GetLocal(emailAddress)
	if err != nil {
		return false, err
	}

	for _, role := range r.roles {
		if local == role {
			return false, nil
		}
	}

	return true, nil
}
