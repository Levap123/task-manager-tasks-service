package mongodb

import (
	"context"
	"github.com/Levap123/task-manager-tasks-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (tr *TaskRepo) GetAll(ctx context.Context, userId int64) ([]models.Task, error) {
	var arr []models.Task
	cur, err := tr.cl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &arr); err != nil {
		return nil, err
	}
	return arr, nil
}
