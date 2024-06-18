package user

import "time"

type User struct {
	id        string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(
	id string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		id:        id,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
