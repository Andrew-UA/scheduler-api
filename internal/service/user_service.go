package service

import (
	"context"
	"scheduler/internal/model"
	"scheduler/internal/repository"
	"scheduler/pkg/logger"
)

type UserService struct {
	Repo   repository.IUser
	Logger logger.Logger
}

func NewUserService(repo repository.IUser, logger logger.Logger) *UserService {
	return &UserService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s UserService) Show(ctx context.Context, ID int) (model.User, error) {
	return s.Repo.FindByID(ID)
}

func (s UserService) Update(ctx context.Context, userID int, m model.User) (model.User, error) {
	return s.Repo.Update(userID, m)
}
