package usecase

import (
	"github.com/yosuke7040/grpc-taskapp/backend/domain/model"
)

type (
	FindTasksByUserIDUseCaseInput struct {
		UserID string
	}

	CreateTaskUseCaseInput struct {
		UserID string
		Name   string
	}

	ChangeTaskNameUseCaseInput struct {
		ID     string
		UserID string
		Name   string
	}

	CompleteTaskUseCaseInput struct {
		ID     string
		UserID string
	}

	UncompleteTaskUseCaseInput struct {
		ID     string
		UserID string
	}

	DeleteTaskUseCaseInput struct {
		ID     string
		UserID string
	}

	FindTasksByUserIDUseCaseOutput struct {
		Tasks []*model.Task
	}
)
