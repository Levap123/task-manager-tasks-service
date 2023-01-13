package repository

import (
	"context"

	"github.com/Levap123/task-manager-tasks-service/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	cl *mongo.Collection
}

// func (tr *TaskRepo) Create(ctx context.Context, in *proto.Task) (*proto.TaskID, error) {

// }
