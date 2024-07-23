package usecase

import (
	"context"

	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
	"github.com/yosuke7040/grpc-taskapp/backend/domain/repository"
)

type IUserUsecase interface {
	FindUserByID(ctx context.Context, in FindUserUseCaseInput) (*model.User, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) IUserUsecase {
	return &userUsecase{
		ur: ur,
	}
}

func (u *userUsecase) FindUserByID(ctx context.Context, in FindUserUseCaseInput) (*model.User, error) {
	id, err := model.NewID(in.ID)
	if err != nil {
		return nil, err
	}

	return u.ur.FindUserByID(ctx, *id)
}
