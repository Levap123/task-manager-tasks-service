package rpc

import (
	"context"

	"github.com/Levap123/task-manager-tasks-service/proto"
)

func (h *Handler) Create(ctx context.Context, in *proto.Task) (*proto.TaskID, error)
