package repository

import (
	"context"

	"github.com/Levap123/task-manager-tasks-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Task
}

type Task interface {
	Create(ctx context.Context, in *proto.Task) (*proto.TaskID, error)
}

func NewRepoMongo(db *mongo.Collection) *Repository {
	return &Repository{}
}

/*
	Create(ctx context.Context, in *Task) (*TaskID, error)
	Update(ctx context.Context, in *Task) (*TaskID, error)
	Get(ctx context.Context, in *TaskID) (*Task, error)
	GetAll(ctx context.Context, in *UserID) (TaskService_GetAllClient, error)
*/
