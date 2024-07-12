package user

import "context"

type UserRepository interface {
	FindUserByID(ctx context.Context, id string) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}
