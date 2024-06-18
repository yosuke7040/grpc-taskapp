// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query_users.sql

package db

import (
	"context"
)

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, email, password, created_at, updated_at
FROM users
WHERE email = ?
LIMIT 1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.queryRow(ctx, q.findUserByEmailStmt, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT id, email, password, created_at, updated_at
FROM users
WHERE id = ?
LIMIT 1
`

func (q *Queries) FindUserByID(ctx context.Context, id string) (*User, error) {
	row := q.queryRow(ctx, q.findUserByIDStmt, findUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
