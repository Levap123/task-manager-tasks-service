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
	GetAll(ctx context.Context, userId int64) ([]models.Task, error)
	Get(ctx context.Context, userId int64, taskId string) (models.Task, error)
	Update(ctx context.Context, title, body, taskId string, userId int64) (string, error)
}

func NewRepoMongo(cl *mongo.Collection) *Repository {
	return &Repository{
		Task: mongodb.NewTaskRepo(cl),
	}
}

/*
	Create(ctx context.Context, in *Task) (*TaskID, error)
	Get(ctx context.Context, in *TaskID) (*Task, error)
*/
