package task

import "context"

type TaskRepository interface {
	FindTaskByID(ctx context.Context, id string) (*Task, error)
	FindTasksByUserID(ctx context.Context, userID string) ([]*Task, error)
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, task *Task) error
	DeleteTask(ctx context.Context, id string) error
}
