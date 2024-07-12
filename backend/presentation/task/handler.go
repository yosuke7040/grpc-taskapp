package task

import (
	"context"

	"connectrpc.com/connect"
	taskApp "github.com/yosuke7040/grpc-taskapp/backend/app/task"
	task_v1 "github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/task/v1"
)

type Handler struct {
	findTaskByIDUseCase      taskApp.FindTaskByIDUseCase
	findTasksByUserIDUseCase taskApp.FindTasksByUserIDUseCase
	createTaskUseCase        taskApp.CreateTaskUseCase
}

func NewHandler(
	findTaskByIDUseCase taskApp.FindTaskByIDUseCase,
	findTasksByUserIDUseCase taskApp.FindTasksByUserIDUseCase,
	createTaskUseCase taskApp.CreateTaskUseCase,
) Handler {
	return Handler{
		findTaskByIDUseCase:      findTaskByIDUseCase,
		findTasksByUserIDUseCase: findTasksByUserIDUseCase,
		createTaskUseCase:        createTaskUseCase,
	}
}

func (h *Handler) taskHandler(
	ctx context.Context,
	arg *connect.Request[task_v1.GetUserRequest],
) task_v1.TaskHandler {
	return task_v1.TaskHandler{
		FindTaskByID:      h.FindTaskByID,
		FindTasksByUserID: h.FindTasksByUserID,
		CreateTask:        h.CreateTask,
	}
}
