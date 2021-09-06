package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"scheduler/internal/http/controller"
	"scheduler/internal/model"
	"scheduler/internal/repository"
	"scheduler/internal/repository/slc"
	"scheduler/internal/server"
	"scheduler/internal/service"
	"scheduler/pkg/auth"
	"scheduler/pkg/router"
	"syscall"
	"time"
)

func Run() {
	// Dependency injection
	// Logger
	loggerService := logrus.New()
	loggerService.SetLevel(logrus.DebugLevel)
	// DB
	pass, _ := auth.NewPasswordManager().HashAndSalt("password")
	var db = &slc.DB{
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
				ID:       1,
				Login:    "Andrew",
				Password: pass,
				Timezone: "Europe/Kiev",
			},
		},
	}
	// Repo
	scheduleRepo := slc.NewSchedules(db)
	userRepo := slc.NewUsers(db)
	r := repository.NewRepository(scheduleRepo, userRepo)

	// ScheduleService
	scheduleService := service.NewScheduleService(r.Schedule, loggerService)
	authService := service.NewJwtService(userRepo, loggerService)
	userService := service.NewUserService(r.User, loggerService)
	s := service.NewService(scheduleService, authService, userService)

	// Middleware
	middleware := &controller.Middleware{
		Logger:   loggerService,
		UserRepo: userRepo,
	}

	// ScheduleController
	c := controller.NewController(
		router.NewRouter(),
		middleware,
		controller.NewScheduleController(s.Schedule, s.User, loggerService),
		controller.NewUserController(userService, loggerService),
		controller.NewAuthController(s.Auth, loggerService),
		controller.NewMetricController(loggerService),
	)
	c.Init()

	srv := server.NewServer(c)

	go func() {
		if err := srv.Run(); err != nil {
			loggerService.Fatal(err)
		}
	}()

	loggerService.Infof("Server started")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithCancel(context.Background())
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		loggerService.Fatalf("Failed to stop HTTP server: %w", err)
	}

	// TODO: DISCONNECT FROM DB
}
