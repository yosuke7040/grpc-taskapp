package usecase

import (
	"context"
	"time"

	"github.com/yosuke7040/grpc-taskapp/backend/app"
	"github.com/yosuke7040/grpc-taskapp/backend/domain"
	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
	"github.com/yosuke7040/grpc-taskapp/backend/domain/repository"
	"github.com/yosuke7040/grpc-taskapp/backend/utils/auth"
	"golang.org/x/crypto/bcrypt"
)

type IAuthUsecase interface {
	Login(ctx context.Context, in *LoginUseCaseInput) (*LoginUseCaseOutput, error)
}

type authUsecase struct {
	ur repository.UserRepository
	auth.ITokenManager
	timeout time.Duration
}

func NewAuthUsecase(ur repository.UserRepository, tm auth.ITokenManager, timeout time.Duration) IAuthUsecase {
	return &authUsecase{
		ur:            ur,
		ITokenManager: tm,
		timeout:       timeout,
	}
}

func (u *authUsecase) Login(ctx context.Context, in *LoginUseCaseInput) (*LoginUseCaseOutput, error) {
	email, err := model.NewEmail(in.Email)
	if err != nil {
		return nil, err
	}

	user, err := u.ur.FindUserByEmail(ctx, *email)
	if err != nil {
		return nil, &domain.ErrNotFound{Msg: "user not found"}
	}

	// bcrypt方式でパスワードが一致するか検証する
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.Value()), []byte(in.Password)); err != nil {
		return nil, &domain.ErrValidationFailed{Msg: "password is incorrect"}
	}

	// JWT作成
	token, err := u.ITokenManager.CreateToken(user.ID().Value(), u.timeout)
	if err != nil {
		return nil, &app.ErrInternal{Msg: "failed to create token"}
	}

	return &LoginUseCaseOutput{
		UserID: user.ID().Value(),
		Email:  user.Email.Value(),
		Token:  token,
	}, nil
}
