package schedule

import (
	"context"
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

func (s Service) List(ctx context.Context, params map[string]string) ([]model.ScheduleEvent, error) {
	return s.Repo.List(ctx, params)
}
func (s Service) Show(ctx context.Context, ID int) (model.ScheduleEvent, error) {
	return s.Repo.Show(ctx, ID)
}
func (s Service) Create(ctx context.Context, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	return s.Repo.Create(ctx, m)
}
func (s Service) Update(ctx context.Context, ID int, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	return s.Repo.Update(ctx, ID, m)
}
func (s Service) Delete(ctx context.Context, ID int) error {
	return s.Repo.Delete(ctx, ID)
}
