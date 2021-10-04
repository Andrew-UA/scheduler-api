package service

import (
	"context"
	"scheduler/internal/helpers"
	"scheduler/internal/model"
	"scheduler/internal/repository"
	"scheduler/pkg/logger"
	"strconv"
)

type ScheduleService struct {
	Repo   repository.ISchedule
	Logger logger.Logger
}

func NewScheduleService(repo repository.ISchedule, logger logger.Logger) *ScheduleService {
	return &ScheduleService{
		Repo:   repo,
		Logger: logger,
	}
}

func (s ScheduleService) List(ctx context.Context, params map[string]string) ([]model.ScheduleEvent, error) {
	authUser, ctxErr := helpers.GetUserFormContext(ctx)
	if ctxErr != nil {
		return nil, ctxErr
	}

	params["user_id"] = strconv.Itoa(authUser.ID)
	return s.Repo.List(params)
}
func (s ScheduleService) Show(ctx context.Context, ID int) (model.ScheduleEvent, error) {
	return s.Repo.Show(ID)
}
func (s ScheduleService) Create(ctx context.Context, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	authUser, ctxErr := helpers.GetUserFormContext(ctx)
	if ctxErr != nil {
		return model.ScheduleEvent{}, ctxErr
	}

	m.UserID = authUser.ID
	return s.Repo.Create(m)
}
func (s ScheduleService) Update(ctx context.Context, ID int, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	return s.Repo.Update(ID, m)
}
func (s ScheduleService) Delete(ctx context.Context, ID int) error {
	return s.Repo.Delete(ID)
}
