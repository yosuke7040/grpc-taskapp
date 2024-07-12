package task

import (
	"context"

	taskDomain "github.com/yosuke7040/grpc-taskapp/backend/domain/model/task"
)

type (
	FindTasksByUserIDUseCase struct {
		taskRepo taskDomain.TaskRepository
	}

	FindTasksByUserIDUseCaseInput struct {
		UserID string
	}

	FindTasksByUserIDUseCaseOutput struct {
		tasks []*taskDomain.Task
	}
)

func NewFindTasksByUserIDUseCase(
	taskRepo taskDomain.TaskRepository,
) *FindTasksByUserIDUseCase {
	return &FindTasksByUserIDUseCase{
		taskRepo: taskRepo,
	}
}

func (u *FindTasksByUserIDUseCase) Execute(ctx context.Context, input FindTasksByUserIDUseCaseInput) (*FindTasksByUserIDUseCaseOutput, error) {
	tasks, err := u.taskRepo.FindTasksByUserID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	return &FindTasksByUserIDUseCaseOutput{
		tasks: tasks,
	}, nil
}
