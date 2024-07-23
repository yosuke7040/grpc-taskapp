package handler

import (
	"context"

	"connectrpc.com/connect"
	"github.com/yosuke7040/grpc-taskapp/backend/app"
	"github.com/yosuke7040/grpc-taskapp/backend/domain"
	task_v1 "github.com/yosuke7040/grpc-taskapp/backend/interfaces/rpc/task/v1"
	"github.com/yosuke7040/grpc-taskapp/backend/usecase"
	"github.com/yosuke7040/grpc-taskapp/backend/utils/contextkey"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskHandler struct {
	taskUsecase   usecase.ITaskUsecase
	contextReader contextkey.IContextReader
}

func NewTaskHandler(
	taskUsecase usecase.ITaskUsecase,
	contextReader contextkey.IContextReader) TaskHandler {
	return TaskHandler{
		taskUsecase:   taskUsecase,
		contextReader: contextReader,
	}
}

func (h *TaskHandler) GetTaskList(ctx context.Context, arg *connect.Request[task_v1.GetTaskListRequest]) (*connect.Response[task_v1.GetTaskListResponse], error) {
	var uid string
	var err error
	if uid, err = h.contextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	res, err := h.taskUsecase.FindTasksByUserID(
		ctx,
		&usecase.FindTasksByUserIDUseCaseInput{UserID: uid},
	)

	if err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	tasks := make([]*task_v1.Task, len(res.Tasks))
	for i, v := range res.Tasks {
		tasks[i] = &task_v1.Task{
			Id:          v.ID().Value(),
			UserId:      v.UserID.Value(),
			Name:        v.Name,
			IsCompleted: v.IsCompleted,
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}
	}
	return connect.NewResponse(&task_v1.GetTaskListResponse{
		Tasks: tasks,
	}), nil
}

func (h *TaskHandler) CreateTask(ctx context.Context, arg *connect.Request[task_v1.CreateTaskRequest]) (*connect.Response[task_v1.CreateTaskResponse], error) {
	var uid string
	var err error
	if uid, err = h.contextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	createdID, err := h.taskUsecase.CreateTask(
		ctx,
		&usecase.CreateTaskUseCaseInput{UserID: uid, Name: arg.Msg.Name},
	)
	if err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.CreateTaskResponse{
		CreatedId: createdID,
	}), nil
}

func (h *TaskHandler) ChangeTaskName(ctx context.Context, arg *connect.Request[task_v1.ChangeTaskNameRequest]) (*connect.Response[task_v1.ChangeTaskNameResponse], error) {
	var uid string
	var err error
	if uid, err = h.contextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := h.taskUsecase.ChangeTaskName(
		ctx,
		&usecase.ChangeTaskNameUseCaseInput{ID: arg.Msg.TaskId, UserID: uid, Name: arg.Msg.Name},
	); err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		case *domain.ErrPermissionDenied:
			return nil, connect.NewError(connect.CodePermissionDenied, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.ChangeTaskNameResponse{}), nil
}

func (h *TaskHandler) CompleteTask(ctx context.Context, arg *connect.Request[task_v1.CompleteTaskRequest]) (*connect.Response[task_v1.CompleteTaskResponse], error) {
	var uid string
	var err error
	if uid, err = h.contextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := h.taskUsecase.CompleteTask(
		ctx,
		&usecase.CompleteTaskUseCaseInput{ID: arg.Msg.TaskId, UserID: uid},
	); err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		case *domain.ErrPermissionDenied:
			return nil, connect.NewError(connect.CodePermissionDenied, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.CompleteTaskResponse{}), nil
}

func (h *TaskHandler) UncompleteTask(ctx context.Context, arg *connect.Request[task_v1.UncompleteTaskRequest]) (*connect.Response[task_v1.UncompleteTaskResponse], error) {
	var uid string
	var err error
	if uid, err = h.contextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := h.taskUsecase.UncompleteTask(
		ctx,
		&usecase.UncompleteTaskUseCaseInput{ID: arg.Msg.TaskId, UserID: uid},
	); err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		case *domain.ErrPermissionDenied:
			return nil, connect.NewError(connect.CodePermissionDenied, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.UncompleteTaskResponse{}), nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, arg *connect.Request[task_v1.DeleteTaskRequest]) (*connect.Response[task_v1.DeleteTaskResponse], error) {
	var uid string
	var err error
	if uid, err = h.contextReader.GetUserID(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, err)
	}

	if err := h.taskUsecase.DeleteTask(
		ctx,
		&usecase.DeleteTaskUseCaseInput{ID: arg.Msg.TaskId, UserID: uid},
	); err != nil {
		switch e := err.(type) {
		case *app.ErrInputValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrValidationFailed:
			return nil, connect.NewError(connect.CodeInvalidArgument, e)
		case *domain.ErrNotFound:
			return nil, connect.NewError(connect.CodeNotFound, e)
		case *domain.ErrPermissionDenied:
			return nil, connect.NewError(connect.CodePermissionDenied, e)
		case *domain.ErrQueryFailed:
			return nil, connect.NewError(connect.CodeAborted, e)
		default:
			return nil, connect.NewError(connect.CodeUnknown, e)
		}
	}
	return connect.NewResponse(&task_v1.DeleteTaskResponse{}), nil

}
