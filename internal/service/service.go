package service

import "scheduler/internal/model"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type IScheduleService interface {
	List(params map[string]string) ([]model.ScheduleEvent, error)
	Show(ID int) (model.ScheduleEvent, error)
	Create(m model.ScheduleEvent) (model.ScheduleEvent, error)
	Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Delete(ID int) error
}

type IAuthService interface {
	SignIn(login, password string) (string, error)
}

type Service struct {
	Schedule IScheduleService
	Auth 	 IAuthService
}

func NewService(schedule IScheduleService, auth IAuthService) *Service {
	return &Service{
		Schedule: schedule,
		Auth: auth,
	}
}