package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
}

type Task interface{
	
}
func NewRepoMongo(db *mongo.Collection) *Repository {
	return &Repository{}
}
