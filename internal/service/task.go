package service

import (
	"context"
	"github.com/Levap123/task-manager-tasks-service/internal/models"
	"github.com/Levap123/task-manager-tasks-service/internal/repository"
	"github.com/Levap123/task-manager-tasks-service/proto"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (ts *TaskService) Create(ctx context.Context, in *proto.Task) (*proto.TaskHelperBody, error) {
	model := &models.Task{
		ID:     "",
		Title:  in.Title,
		Body:   in.Body,
		UserID: in.UserId,
	}
	taskId, err := ts.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}
	modelResp := &proto.TaskHelperBody{
		Id: taskId,
	}
	return modelResp, err
}
