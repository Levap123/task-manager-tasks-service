package repository

import (
	"context"
	"github.com/Levap123/task-manager-tasks-service/internal/models"
	"github.com/Levap123/task-manager-tasks-service/internal/repository/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Task
}

type Task interface {
	Create(ctx context.Context, in *models.Task) (string, error)
}

func NewRepoMongo(cl *mongo.Collection) *Repository {
	return &Repository{
		Task: mongodb.NewTaskRepo(cl),
	}
}

/*
	Create(ctx context.Context, in *Task) (*TaskID, error)
	Update(ctx context.Context, in *Task) (*TaskID, error)
	Get(ctx context.Context, in *TaskID) (*Task, error)
	GetAll(ctx context.Context, in *UserID) (TaskService_GetAllClient, error)
*/
