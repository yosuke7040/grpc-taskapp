package user

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yosuke7040/grpc-taskapp/backend/app"
	userApp "github.com/yosuke7040/grpc-taskapp/backend/app/user"
	user_v1 "github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/user/v1"
)

type Handler struct {
	findUserUseCase *userApp.FindUserUseCase
}

func NewHandler(
	findUserUseCase *userApp.FindUserUseCase,
) Handler {
	return Handler{
		findUserUseCase: findUserUseCase,
	}
}

func (h *Handler) GetUser(
	ctx context.Context,
	arg *connect.Request[user_v1.GetUserRequest],
) (
	*connect.Response[user_v1.GetUserResponse],
	error,
) {
	// ? validationってどこでするのが良いんだろうか?domain serviceみたいなところ作る？
	var input userApp.FindUserUseCaseInput
	input.ID = arg.Msg.Id

	user, err := h.findUserUseCase.Execute(ctx, input)
	if err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		// case *domain.ErrInputValidationFailed:
		// 	return nil, connect.NewError(connect.CodeInvalidArgument, e)
		// case *domain.ErrUserNotFound:
		// 	return nil, connect.NewError(connect.CodeNotFound, e)
		default:
			return nil, connect.NewError(connect.CodeInternal, e)
		}
	}

	return connect.NewResponse(&user_v1.GetUserResponse{
		User: &user_v1.User{
			Id:    user.ID,
			Email: user.Email,
		},
	}), nil
}
