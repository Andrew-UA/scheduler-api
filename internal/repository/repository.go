package repository

import (
	"context"
	"scheduler/internal/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type ISchedule interface {
	List(ctx context.Context, params map[string]string) ([]model.ScheduleEvent, error)
	Show(ctx context.Context, D int) (model.ScheduleEvent, error)
	Create(ctx context.Context, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Update(ctx context.Context, ID int, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Delete(ctx context.Context, ID int) error
}

type IUser interface {
	FindByLogin(login string) (model.User, error)
	FindByID(ID int) (model.User, error)
	Update(userID int, user model.User) (model.User, error)
}

type Repository struct {
	Schedule ISchedule
	User     IUser
}

func NewRepository(schedule ISchedule, user IUser) *Repository {
	return &Repository{
		Schedule: schedule,
		User:     user,
	}
}
