package model

import "time"

type User struct {
	id        ID
	Email     Email
	Password  Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(
	id ID,
	email Email,
	password Password,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	return &User{
		id:        id,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (u *User) ID() ID {
	return u.id
}
