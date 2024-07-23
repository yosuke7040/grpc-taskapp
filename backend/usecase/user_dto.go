package usecase

import (
	"time"
)

type (
	FindUserUseCaseInput struct {
		ID string
	}

	FindUserUseCaseOutput struct {
		ID        string
		Email     string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
