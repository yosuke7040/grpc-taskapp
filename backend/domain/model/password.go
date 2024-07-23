package model

import "github.com/yosuke7040/grpc-taskapp/backend/domain"

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	if value == "" {
		return nil, &domain.ErrValidationFailed{Msg: "Password is required"}
	}

	return &Password{value: value}, nil
}

func (p *Password) Value() string {
	return p.value
}
