package repository

import (
	"context"
	"fmt"

	model "github.com/yosuke7040/grpc-taskapp/backend/domain/model"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/model/db"
)

type TaskRepository struct {
	db.Querier
}

func NewTaskRepository(qry db.Querier) *TaskRepository {
	return &TaskRepository{qry}
}

func (r *TaskRepository) FindTaskByID(ctx context.Context, id string) (*model.Task, error) {
	res, err := r.Querier.FindTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	taskID, err := model.NewID(res.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID from database: %w", err)
	}

	userID, err := model.NewID(res.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID from database: %w", err)
	}

	return model.NewTask(
		*taskID,
		*userID,
		res.Name,
		res.IsCompleted,
		res.CreatedAt,
		res.UpdatedAt,
	), nil
}

func (r *TaskRepository) FindTasksByUserID(ctx context.Context, userID string) ([]*model.Task, error) {
	res, err := r.Querier.FindTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	tasks := make([]*model.Task, len(res))
	for i, v := range res {
		taskID, err := model.NewID(v.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid user ID from database: %w", err)
		}

		userID, err := model.NewID(v.UserID)
		if err != nil {
			return nil, fmt.Errorf("invalid user ID from database: %w", err)
		}
		tasks[i] = model.NewTask(
			*taskID,
			*userID,
			v.Name,
			v.IsCompleted,
			v.CreatedAt,
			v.UpdatedAt,
		)
	}

	return tasks, nil
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	return r.Querier.CreateTask(ctx, db.CreateTaskParams{
		// 値オブジェクトやプライベートフィールドに対するgetter的なやつがあったりごちゃついてる様にも見える…
		ID:          task.ID().Value(),
		UserID:      task.UserID.Value(),
		Name:        task.Name,
		IsCompleted: task.IsCompleted,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	})
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	return r.Querier.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          task.ID().Value(),
		Name:        task.Name,
		IsCompleted: task.IsCompleted,
		UpdatedAt:   task.UpdatedAt,
	})
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id string) error {
	return r.Querier.DeleteTask(ctx, id)
}
