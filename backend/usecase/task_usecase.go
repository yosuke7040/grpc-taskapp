package usecase

import (
	"context"
	"html"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/yosuke7040/grpc-taskapp/backend/domain"
	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
	"github.com/yosuke7040/grpc-taskapp/backend/domain/repository"
)

type ITaskUsecase interface {
	FindTasksByUserID(ctx context.Context, in *FindTasksByUserIDUseCaseInput) (*FindTasksByUserIDUseCaseOutput, error)
	CreateTask(ctx context.Context, in *CreateTaskUseCaseInput) (string, error)
	ChangeTaskName(ctx context.Context, in *ChangeTaskNameUseCaseInput) error
	CompleteTask(ctx context.Context, id *CompleteTaskUseCaseInput) error
	UncompleteTask(ctx context.Context, in *UncompleteTaskUseCaseInput) error
	DeleteTask(ctx context.Context, in *DeleteTaskUseCaseInput) error
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func NewTaskUsecase(tr repository.TaskRepository) ITaskUsecase {
	return &taskUsecase{
		tr: tr,
	}
}

func (u *taskUsecase) FindTasksByUserID(ctx context.Context, in *FindTasksByUserIDUseCaseInput) (*FindTasksByUserIDUseCaseOutput, error) {
	userID, err := model.NewID(in.UserID)
	if err != nil {
		return nil, err
	}

	tasks, err := u.tr.FindTasksByUserID(ctx, userID.Value())
	if err != nil {
		return nil, err
	}

	return &FindTasksByUserIDUseCaseOutput{
		Tasks: tasks,
	}, nil
}

func (u *taskUsecase) CreateTask(ctx context.Context, in *CreateTaskUseCaseInput) (string, error) {
	// TODO: 多分ドメインサービスとかになると思う
	now := time.Now()
	slog.Info("Task creation time", "now", now)
	userID, err := model.NewID(in.UserID)
	if err != nil {
		return "", err
	}
	id, err := model.NewID(uuid.NewString())
	if err != nil {
		return "", err
	}

	task := model.NewTask(
		*id, *userID, html.EscapeString(in.Name), false, now, now,
	)

	slog.Info("Creating task", "created_at", task.CreatedAt, "updated_at", task.UpdatedAt)

	err = u.tr.CreateTask(ctx, task)
	if err != nil {
		return "", err
	}

	return id.Value(), nil
}

func (u *taskUsecase) ChangeTaskName(ctx context.Context, in *ChangeTaskNameUseCaseInput) error {
	taskID, err := model.NewID(in.ID)
	if err != nil {
		return err
	}
	userID, err := model.NewID(in.UserID)
	if err != nil {
		return err
	}

	task, err := u.tr.FindTaskByID(ctx, taskID.Value())
	if err != nil {
		return err
	}
	if !task.UserID.Equals(userID) {
		return &domain.ErrPermissionDenied{}
	}

	task.ChangeName(html.EscapeString(in.Name))
	task.UpdateUpdatedAt(time.Now())

	return u.tr.UpdateTask(ctx, task)
}

func (u *taskUsecase) CompleteTask(ctx context.Context, in *CompleteTaskUseCaseInput) error {
	taskID, err := model.NewID(in.ID)
	if err != nil {
		return err
	}
	userID, err := model.NewID(in.UserID)
	if err != nil {
		return err
	}

	task, err := u.tr.FindTaskByID(ctx, taskID.Value())
	if err != nil {
		return err
	}
	if !task.UserID.Equals(userID) {
		return &domain.ErrPermissionDenied{}
	}
	task.Complete()
	task.UpdateUpdatedAt(time.Now())

	return u.tr.UpdateTask(ctx, task)
}

func (u *taskUsecase) UncompleteTask(ctx context.Context, in *UncompleteTaskUseCaseInput) error {
	taskID, err := model.NewID(in.ID)
	if err != nil {
		return err
	}
	userID, err := model.NewID(in.UserID)
	if err != nil {
		return err
	}

	task, err := u.tr.FindTaskByID(ctx, taskID.Value())
	if err != nil {
		return err
	}
	if !task.UserID.Equals(userID) {
		return &domain.ErrPermissionDenied{}
	}
	task.Uncomplete()
	task.UpdateUpdatedAt(time.Now())

	return u.tr.UpdateTask(ctx, task)
}

func (u *taskUsecase) DeleteTask(ctx context.Context, in *DeleteTaskUseCaseInput) error {
	taskID, err := model.NewID(in.ID)
	if err != nil {
		return err
	}
	userID, err := model.NewID(in.UserID)
	if err != nil {
		return err
	}

	task, err := u.tr.FindTaskByID(ctx, taskID.Value())
	if err != nil {
		return err
	}
	if !task.UserID.Equals(userID) {
		return &domain.ErrPermissionDenied{}
	}

	return u.tr.DeleteTask(ctx, taskID.Value())
}
