package mongodb

import (
	"context"
	"github.com/Levap123/task-manager-tasks-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	cl *mongo.Collection
}

func NewTaskRepo(cl *mongo.Collection) *TaskRepo {
	return &TaskRepo{cl: cl}
}
func (tr *TaskRepo) Create(ctx context.Context, in *models.Task) (string, error) {
	res, err := tr.cl.InsertOne(ctx, in)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
