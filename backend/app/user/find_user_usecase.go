package user

import (
	"context"
	"time"

	userDomain "github.com/yosuke7040/grpc-taskapp/backend/domain/model/user"
)

type (
	FindUserUseCase struct {
		userRepo userDomain.UserRepository
	}

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

func NewFindUserUseCase(
	userRepo userDomain.UserRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

func (u *FindUserUseCase) Execute(ctx context.Context, input FindUserUseCaseInput) (*FindUserUseCaseOutput, error) {
	user, err := u.userRepo.FindUserByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &FindUserUseCaseOutput{
		ID:        user.ID(),
		Email:     user.Email(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}
