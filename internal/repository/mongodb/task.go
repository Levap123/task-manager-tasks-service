package mongodb

import (
	"context"
	"fmt"
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
	cur, err := tr.cl.Find(ctx, bson.D{{"user_id", userId}})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &arr); err != nil {
		return nil, err
	}
	return arr, nil
}
func (tr *TaskRepo) Get(ctx context.Context, userId int64, taskId string) (models.Task, error) {
	var task models.Task
	objectId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return models.Task{}, err
	}

	if err := tr.cl.FindOne(ctx, bson.D{{"user_id", userId}, {"_id", objectId}}).Decode(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (tr *TaskRepo) Update(ctx context.Context, title, body, taskId string, userId int64) (string, error) {
	updateBody := bson.D{
		{
			"title", title,
		},
		{
			"body", body,
		},
		{
			"user_id", userId,
		},
	}
	update := bson.D{{"$set", updateBody}}
	taskIdObj, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return "", err
	}
	res, err := tr.cl.UpdateByID(ctx, taskIdObj, update)
	if res.MatchedCount == 0 {
		return "", fmt.Errorf("taskId not found")
	}
	if err != nil {
		return "", err
	}
	return taskId, nil
}
