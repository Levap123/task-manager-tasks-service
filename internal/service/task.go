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

func (ts *TaskService) GetAll(ctx context.Context, in *proto.UserRequest) (*proto.TaskArray, error) {
	res, err := ts.repo.GetAll(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	protoArr := make([]*proto.Task, 0)
	for i := 0; i < len(res); i++ {
		buff := &proto.Task{
			Id:     res[i].ID,
			Title:  res[i].Title,
			Body:   res[i].Body,
			UserId: res[i].UserID,
		}
		protoArr = append(protoArr, buff)
	}
	req := &proto.TaskArray{
		Tasks: protoArr,
	}
	return req, nil
}

func (ts *TaskService) Get(ctx context.Context, in *proto.UserAndTask) (*proto.Task, error) {
	task, err := ts.repo.Get(ctx, in.UserId, in.TaskId)
	if err != nil {
		return nil, err
	}
	resp := &proto.Task{
		Id:     task.ID,
		Title:  task.Title,
		Body:   task.Body,
		UserId: task.UserID,
	}
	return resp, nil
}

func (ts *TaskService) Update(ctx context.Context, task *proto.Task) (*proto.TaskHelperBody, error) {
	taskId, err := ts.repo.Update(ctx, task.Title, task.Body, task.Id, task.UserId)
	if err != nil {
		return nil, err
	}
	resp := &proto.TaskHelperBody{
		Id: taskId,
	}
	return resp, nil
}
