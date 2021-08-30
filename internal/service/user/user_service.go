package user

import (
	"context"
	"scheduler/internal/model"
	"scheduler/internal/repository"
)

type Service struct {
	Repo repository.IUser
}

func NewService(repo repository.IUser) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s Service) Show(ctx context.Context, ID int) (model.User, error)  {
	return s.Repo.FindByID(ID)
}

func (s Service) Update(ctx context.Context, userID int, m model.User) (model.User, error) {
	return s.Repo.Update(userID, m)
}
