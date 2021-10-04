package repository

import (
	"scheduler/internal/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type ISchedule interface {
	List(params map[string]string) ([]model.ScheduleEvent, error)
	Show(ID int) (model.ScheduleEvent, error)
	Create(m model.ScheduleEvent) (model.ScheduleEvent, error)
	Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Delete(ID int) error
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
