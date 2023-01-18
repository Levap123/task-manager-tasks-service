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

func (h *Handler) Get(ctx context.Context, req *proto.UserAndTask) (*proto.Task, error) {
	return h.service.Get(ctx, req)
}

func (h *Handler) Update(ctx context.Context, req *proto.Task) (*proto.TaskHelperBody, error) {
	return h.service.Update(ctx, req)
}

/*
	Create(context.Context, *Task) (*TaskHelperBody, error)
	Update(context.Context, *Task) (*TaskHelperBody, error)
	GetAll(context.Context, *UserRequest) (*TaskArray, error)
	Delete(context.Context, *TaskHelperBody) (*TaskHelperBody, error)
*/
