package task

import (
	"context"
	"time"

	taskDomain "github.com/yosuke7040/grpc-taskapp/backend/domain/model/task"
)

type (
	CreateTaskUseCase struct {
		taskRepo taskDomain.TaskRepository
	}

	CreateTaskUseCaseInput struct {
		ID          string
		UserID      string
		Name        string
		IsCompleted bool
	}

	CreateTaskUseCaseOutput struct {
	}
)

func NewCreateTaskUseCase(
	taskRepo taskDomain.TaskRepository,
) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		taskRepo: taskRepo,
	}
}

func (u *CreateTaskUseCase) Execute(ctx context.Context, input CreateTaskUseCaseInput) (*CreateTaskUseCaseOutput, error) {
	now := time.Now().UTC()
	err := u.taskRepo.CreateTask(ctx, taskDomain.NewTask(
		input.ID,
		input.UserID,
		input.Name,
		input.IsCompleted,
		now, // createdAt
		now, // updatedAt
	))

	if err != nil {
		return nil, err
	}

	return &CreateTaskUseCaseOutput{}, nil
}
