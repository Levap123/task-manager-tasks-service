package service

import "github.com/Levap123/task-manager-tasks-service/internal/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
