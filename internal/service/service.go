package service

import (
	"context"
	"scheduler/internal/model"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type IScheduleService interface {
	List(ctx context.Context, params map[string]string) ([]model.ScheduleEvent, error)
	Show(ctx context.Context, ID int) (model.ScheduleEvent, error)
	Create(ctx context.Context, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Update(ctx context.Context, ID int, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Delete(ctx context.Context, ID int) error
}

type IAuthService interface {
	SignIn(login, password string) (string, error)
}

type IUserService interface {
	Show(ctx context.Context, ID int) (model.User, error)
	Update(ctx context.Context, userID int, m model.User) (model.User, error)
}

type Service struct {
	Schedule IScheduleService
	Auth 	 IAuthService
	User     IUserService
}

func NewService(schedule IScheduleService, auth IAuthService, user IUserService) *Service {
	return &Service{
		Schedule: schedule,
		Auth: auth,
		User: user,
	}
}