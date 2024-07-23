package repository

import (
	"context"

	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
)

type TaskRepository interface {
	FindTaskByID(ctx context.Context, id string) (*model.Task, error)
	FindTasksByUserID(ctx context.Context, userID string) ([]*model.Task, error)
	CreateTask(ctx context.Context, task *model.Task) error
	UpdateTask(ctx context.Context, task *model.Task) error
	DeleteTask(ctx context.Context, id string) error
}
