package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yosuke7040/grpc-taskapp/backend/app"
	"github.com/yosuke7040/grpc-taskapp/backend/domain"
	user_v1 "github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/user/v1"
	"github.com/yosuke7040/grpc-taskapp/backend/usecase"
)

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(
	userUsecase usecase.IUserUsecase,
) UserHandler {
	return UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetUser(
	ctx context.Context,
	arg *connect.Request[user_v1.GetUserRequest],
) (
	*connect.Response[user_v1.GetUserResponse],
	error,
) {
	// ここでチェックするのは基本的な入力チェックや、形式的なバリデーションのみ
	// なので、IDはstringであることだけチェックが良さそう。ユースケースで実際にビジネスルールにそった条件下チェックするのが良さそう？
	// ここのdtoはmodelに依存せず、プリミティブ型で指定した
	input := usecase.FindUserUseCaseInput{
		ID: arg.Msg.Id,
	}

	user, err := h.userUsecase.FindUserByID(ctx, input)
	if err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		default:
			return nil, connect.NewError(connect.CodeInternal, e)
		}
	}

	return connect.NewResponse(&user_v1.GetUserResponse{
		User: &user_v1.User{
			// IDはプライベートな値にしたけど、う〜〜〜ん
			Id:    user.ID().Value(),
			Email: user.Email.Value(),
		},
	}), nil
}
