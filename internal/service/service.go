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
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Task: NewTaskService(repo.Task),
	}
}
