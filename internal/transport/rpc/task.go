package rpc

import (
	"context"

	"github.com/Levap123/task-manager-tasks-service/proto"
)

func (h *Handler) Create(ctx context.Context, in *proto.Task) (*proto.TaskHelperBody, error) {
	return h.service.Create(ctx, in)
}

func (h *Handler) GetAll(ctx context.Context, req *proto.UserRequest) (*proto.TaskArray, error) {
	return h.service.GetAll(ctx, req)
}

/*
	Create(context.Context, *Task) (*TaskHelperBody, error)
	Update(context.Context, *Task) (*TaskHelperBody, error)
	GetAll(context.Context, *UserRequest) (*TaskArray, error)
	Get(context.Context, *UserAndTask) (*Task, error)
	Delete(context.Context, *TaskHelperBody) (*TaskHelperBody, error)
*/
