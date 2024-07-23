package model

import "github.com/yosuke7040/grpc-taskapp/backend/domain"

type ID struct {
	value string
}

func NewID(value string) (*ID, error) {
	if value == "" {
		return nil, &domain.ErrValidationFailed{Msg: "ID is required"}
	}

	// TODO: validation

	return &ID{value: value}, nil
}

func (id ID) Value() string {
	return id.value
}

func (id ID) Equals(other *ID) bool {
	return id.value == other.value
}
