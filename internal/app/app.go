package app

import (
	"scheduler/internal/http/controller"
	controllerauth "scheduler/internal/http/controller/auth"
	"scheduler/internal/http/controller/schedule"
	"scheduler/internal/http/controller/user"
	"scheduler/internal/model"
	"scheduler/internal/repository"
	"scheduler/internal/repository/sl"
	"scheduler/internal/server"
	"scheduler/internal/service"
	serviceAuth "scheduler/internal/service/auth"
	serviceSchedule "scheduler/internal/service/schedule"
	serviceUser "scheduler/internal/service/user"
	"scheduler/pkg/auth"
	"scheduler/pkg/router"
	"time"
)

func Run() error {
	// Dependency injection
	// DB
	pass, _ := auth.NewPasswordManager().HashAndSalt("password")
	var db = &sl.DB{
		ScheduleIncrement: 2,
		Schedules: []model.ScheduleEvent{
			{
				ID:        1,
				UserID:    2,
				Name:      "First",
				Time:      160,
				StartAt:   time.Now().Add(24 * time.Hour).Unix(),
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
			{
				ID:        2,
				UserID:    1,
				Name:      "Second",
				Time:      160,
				StartAt:   time.Now().Add(48 * time.Hour).Unix(),
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			},
		},
		Users: []model.User{
			{
				ID: 1,
				Login: "Andrew",
				Password: pass,
				Timezone: "Europe/Kiev",
			},
		},
	}
	// Repo
	scheduleRepo := sl.NewSchedules(db)
	userRepo := sl.NewUsers(db)
	r := repository.NewRepository(scheduleRepo, userRepo)

	// Service
	scheduleService := serviceSchedule.NewService(r.Schedule)
	authService := serviceAuth.NewJwtService(userRepo)
	userService := serviceUser.NewService(r.User)
	s := service.NewService(scheduleService, authService, userService)

	// Controller
	c := controller.NewController(
		router.NewRouter(),
		schedule.NewController(s.Schedule, s.User),
		user.NewController(userService),
		controllerauth.NewController(s.Auth),
	)
	c.Init()

	srv := server.NewServer(c)

	return srv.Run()
}
