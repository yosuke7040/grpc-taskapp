package repository

import (
	"context"
	"fmt"

	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/model/db"
)

type UserRepository struct {
	db.Querier
}

func NewUserRepository(qry db.Querier) *UserRepository {
	return &UserRepository{qry}
}

func (r *UserRepository) FindUserByID(ctx context.Context, userID model.ID) (*model.User, error) {
	res, err := r.Querier.FindUserByID(ctx, userID.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	// TODO:これ以降はほかメソッドとも重複するから、ファクトリメソッドとかバリデーションメソッドにしてしまう？
	id, err := model.NewID(res.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID from database: %w", err)
	}

	email, err := model.NewEmail(res.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email from database: %w", err)
	}

	password, err := model.NewPassword(res.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid password from database: %w", err)
	}

	user, err := model.NewUser(
		*id,
		*email,
		*password,
		res.CreatedAt,
		res.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user model: %w", err)
	}

	return user, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email model.Email) (*model.User, error) {
	res, err := r.Querier.FindUserByEmail(ctx, email.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to query user: %w", err)
	}

	id, err := model.NewID(res.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID from database: %w", err)
	}

	emailFromDB, err := model.NewEmail(res.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email from database: %w", err)
	}

	password, err := model.NewPassword(res.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid password from database: %w", err)
	}

	user, err := model.NewUser(
		*id,
		*emailFromDB,
		*password,
		res.CreatedAt,
		res.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user model: %w", err)
	}

	return user, nil
}
