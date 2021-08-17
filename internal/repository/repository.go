package repository

import "scheduler/internal/model"

type ISchedule interface {
	List(params map[string]string) ([]model.ScheduleEvent, error)
	Show(ID int) (model.ScheduleEvent, error)
	Create(m model.ScheduleEvent) (model.ScheduleEvent, error)
	Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error)
	Delete(ID int) error
}

type Repository struct {
	Schedule ISchedule
}

func NewRepository(schedule ISchedule) *Repository {
	return &Repository{
		Schedule: schedule,
	}
}