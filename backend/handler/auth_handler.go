package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yosuke7040/grpc-taskapp/backend/app"
	"github.com/yosuke7040/grpc-taskapp/backend/domain"
	"github.com/yosuke7040/grpc-taskapp/backend/usecase"

	auth_v1 "github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/auth/v1"
)

type AuthHandler struct {
	authUsecase usecase.IAuthUsecase
}

func NewAuthHandler(
	authUsecase usecase.IAuthUsecase,
) AuthHandler {
	return AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Login(ctx context.Context, arg *connect.Request[auth_v1.LoginRequest]) (*connect.Response[auth_v1.LoginResponse], error) {
	res, err := h.authUsecase.Login(
		ctx,
		&usecase.LoginUseCaseInput{
			Email:    arg.Msg.Email,
			Password: arg.Msg.Password,
		},
	)

	if err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *app.ErrLoginFailed:
			return nil, connect.NewError(connect.CodeUnauthenticated, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeUnauthenticated, e)
		case *app.ErrInternal:
			return nil, connect.NewError(connect.CodeInternal, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&auth_v1.LoginResponse{
		Token: res.Token,
	}), nil
}
