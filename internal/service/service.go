package service

import "scheduler/internal/model"

type IScheduleService interface {
	List(params map[string]string) ([]model.ScheduleEvent, error)
	Show(ID int) (model.ScheduleEvent, error)
	Create(m model.ScheduleEvent) (model.ScheduleEvent, error)
	Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Delete(ID int) error
}

type Service struct {
	Schedule IScheduleService
}

func NewService(schedule IScheduleService) *Service {
	return &Service{
		Schedule: schedule,
	}
}