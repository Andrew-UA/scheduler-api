package app

import (
	"scheduler/internal/http"
	httpSchedule "scheduler/internal/http/schedule"
	"scheduler/internal/model"
	"scheduler/internal/repository"
	"scheduler/internal/repository/sl"
	"scheduler/internal/server"
	"scheduler/internal/service"
	serviceSchedule "scheduler/internal/service/schedule"
	"scheduler/pkg/router"
	"time"
)

func Run() error {
	// Dependency injection
	// DB
	var db = &sl.DB{
		ScheduleIncrement: 2,
		Schedules: []model.ScheduleEvent{
			{
				ID:        1,
				Name:      "First",
				Time:      160,
				StartAt:   time.Now().Add(24 * time.Hour).Unix(),
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
			{
				ID:        2,
				Name:      "Second",
				Time:      160,
				StartAt:   time.Now().Add(48 * time.Hour).Unix(),
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
	// Repo
	scheduleRepo := sl.NewSchedules(db)
	r := repository.NewRepository(scheduleRepo)

	// Service
	scheduleService := serviceSchedule.NewService(r.Schedule)
	s := service.NewService(scheduleService)

	// Controller
	c := http.NewController(
		router.NewRouter(),
		httpSchedule.NewController(s.Schedule),
	)
	c.Init()

	srv := server.NewServer(c)

	return srv.Run()
}
