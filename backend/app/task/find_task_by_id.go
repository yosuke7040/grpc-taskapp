package task

import (
	"context"

	taskDomain "github.com/yosuke7040/grpc-taskapp/backend/domain/model/task"
)

type (
	FindTaskByIDUseCase struct {
		taskRepo taskDomain.TaskRepository
	}

	FindTaskByIDUseCaseInput struct {
		ID string
	}

	FindTaskByIDUseCaseOutput struct {
		Task *taskDomain.Task
	}
)

func NewFindTaskByIDUseCase(
	taskRepo taskDomain.TaskRepository,
) *FindTaskByIDUseCase {
	return &FindTaskByIDUseCase{
		taskRepo: taskRepo,
	}
}

func (u *FindTaskByIDUseCase) Execute(ctx context.Context, input FindTaskByIDUseCaseInput) (*FindTaskByIDUseCaseOutput, error) {
	task, err := u.taskRepo.FindTaskByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &FindTaskByIDUseCaseOutput{
		Task: task,
	}, nil
}
