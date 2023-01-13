package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDb() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTasksCollection(cl *mongo.Client) *mongo.Collection {
	return cl.Database("service").Collection("tasks")
}
