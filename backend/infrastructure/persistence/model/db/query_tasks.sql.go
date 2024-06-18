// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query_tasks.sql

package db

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :execresult
INSERT INTO tasks(id, user_id, name, is_completed, created_at, updated_at)
VALUES(?, ?, ?, ?, ?, ?)
`

type CreateTaskParams struct {
	ID          string       `json:"id"`
	UserID      string       `json:"user_id"`
	Name        string       `json:"name"`
	IsCompleted bool         `json:"is_completed"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (sql.Result, error) {
	return q.exec(ctx, q.createTaskStmt, createTask,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.IsCompleted,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?
`

func (q *Queries) DeleteTask(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteTaskStmt, deleteTask, id)
	return err
}

const findTaskByID = `-- name: FindTaskByID :one
SELECT id, user_id, name, is_completed, created_at, updated_at
FROM tasks
WHERE id = ?
LIMIT 1
`

func (q *Queries) FindTaskByID(ctx context.Context, id string) (*Task, error) {
	row := q.queryRow(ctx, q.findTaskByIDStmt, findTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.IsCompleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const findTasksByUserID = `-- name: FindTasksByUserID :many
SELECT id, user_id, name, is_completed, created_at, updated_at
FROM tasks
WHERE user_id = ?
ORDER BY updated_at DESC
`

func (q *Queries) FindTasksByUserID(ctx context.Context, userID string) ([]*Task, error) {
	rows, err := q.query(ctx, q.findTasksByUserIDStmt, findTasksByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.IsCompleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks
SET name = ?, is_completed = ?, updated_at = ?
WHERE id = ?
`

type UpdateTaskParams struct {
	Name        string       `json:"name"`
	IsCompleted bool         `json:"is_completed"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	ID          string       `json:"id"`
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.exec(ctx, q.updateTaskStmt, updateTask,
		arg.Name,
		arg.IsCompleted,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
