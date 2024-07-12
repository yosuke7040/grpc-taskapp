package repository

import (
	"context"

	taskDomain "github.com/yosuke7040/grpc-taskapp/backend/domain/model/task"
	"github.com/yosuke7040/grpc-taskapp/backend/infrastructure/persistence/model/db"
)

type TaskRepository struct {
	db.Querier
}

func NewTaskRepository(qry db.Querier) *TaskRepository {
	return &TaskRepository{qry}
}

func (r *TaskRepository) FindTaskByID(ctx context.Context, id string) (*taskDomain.Task, error) {
	res, err := r.Querier.FindTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return taskDomain.NewTask(
		res.ID,
		res.UserID,
		res.Name,
		res.IsCompleted,
		res.CreatedAt,
		res.UpdatedAt,
	), nil
}

func (r *TaskRepository) FindTasksByUserID(ctx context.Context, userID string) ([]*taskDomain.Task, error) {
	res, err := r.Querier.FindTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	tasks := make([]*taskDomain.Task, len(res))
	for i, v := range res {
		tasks[i] = taskDomain.NewTask(
			v.ID,
			v.UserID,
			v.Name,
			v.IsCompleted,
			v.CreatedAt,
			v.UpdatedAt,
		)
	}

	return tasks, nil
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *taskDomain.Task) error {
	return r.Querier.CreateTask(ctx, db.CreateTaskParams{
		ID:          task.ID(),
		UserID:      task.UserID(),
		Name:        task.Name(),
		IsCompleted: task.IsCompleted(),
		CreatedAt:   task.CreatedAt(),
	})
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *taskDomain.Task) error {
	return r.Querier.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          task.ID(),
		Name:        task.Name(),
		IsCompleted: task.IsCompleted(),
	})
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id string) error {
	return r.Querier.DeleteTask(ctx, id)
}
