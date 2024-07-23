package repository

import (
	"context"

	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id model.ID) (*model.User, error)
	FindUserByEmail(ctx context.Context, email model.Email) (*model.User, error)
}
