package service

import (
	"context"
	"github.com/Levap123/task-manager-tasks-service/internal/repository"
	"github.com/Levap123/task-manager-tasks-service/proto"
)

type Service struct {
	Task
}

type Task interface {
	Create(ctx context.Context, in *proto.Task) (*proto.TaskHelperBody, error)
	GetAll(ctx context.Context, in *proto.UserRequest) (*proto.TaskArray, error)
	Get(ctx context.Context, in *proto.UserAndTask) (*proto.Task, error)
	Update(ctx context.Context, in *proto.Task) (*proto.TaskHelperBody, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(repo.Task),
	}
}
