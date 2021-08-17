package schedule

import (
	"scheduler/internal/model"
	"scheduler/internal/repository"
)

type Service struct {
	Repo repository.ISchedule
}

func NewService(repo repository.ISchedule) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s Service) List(params map[string]string) ([]model.ScheduleEvent, error) {
	return s.Repo.List(params)
}
func (s Service) Show(ID int) (model.ScheduleEvent, error) {
	return s.Repo.Show(ID)
}
func (s Service) Create(m model.ScheduleEvent) (model.ScheduleEvent, error) {
	return s.Repo.Create(m)
}
func (s Service) Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	return s.Repo.Update(ID, m)
}
func (s Service) Delete(ID int) error {
	return s.Repo.Delete(ID)
}
