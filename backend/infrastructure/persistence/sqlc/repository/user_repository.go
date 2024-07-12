package repository

import (
	"context"

	userDomain "github.com/yosuke7040/grpc-taskapp/backend/domain/model/user"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/model/db"
)

type UserRepository struct {
	db.Querier
}

func NewUserRepository(qry db.Querier) *UserRepository {
	return &UserRepository{qry}
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*userDomain.User, error) {
	res, err := r.Querier.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// var createdAt time.Time
	// if res.CreatedAt.Valid {
	// 	createdAt = res.CreatedAt.Time
	// }

	// var updatedAt time.Time
	// if res.UpdatedAt.Valid {
	// 	updatedAt = res.UpdatedAt.Time
	// }

	return userDomain.NewUser(
		res.ID,
		res.Email,
		res.Password,
		res.CreatedAt,
		res.UpdatedAt,
		// createdAt,
		// updatedAt,
	), nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, id string) (*userDomain.User, error) {
	res, err := r.Querier.FindUserByEmail(ctx, id)
	if err != nil {
		return nil, err
	}

	// var createdAt time.Time
	// if res.CreatedAt.Valid {
	// 	createdAt = res.CreatedAt.Time
	// }

	// var updatedAt time.Time
	// if res.UpdatedAt.Valid {
	// 	updatedAt = res.UpdatedAt.Time
	// }

	return userDomain.NewUser(
		res.ID,
		res.Email,
		res.Password,
		res.CreatedAt,
		res.UpdatedAt,
	), nil
}
