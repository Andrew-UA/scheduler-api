package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"scheduler/internal/http/controller"
	"scheduler/internal/repository"
	"scheduler/internal/repository/postgres"
	"scheduler/internal/server"
	"scheduler/internal/service"
	"scheduler/pkg/router"
	"syscall"
)

// Run initialize application
func Run() {
	// Dependency injection
	// Logger
	loggerService := logrus.New()
	loggerService.SetLevel(logrus.DebugLevel)

	/*if err := initConfig(); err != nil {
		loggerService.Fatalf("Error init config: %s", err.Error())
	}*/

	// DB
	/*pass, _ := auth.NewPasswordManager().HashAndSalt("password")
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
	}*/
	// Repo
	//scheduleRepo := slc.NewSchedules(db)
	//userRepo := slc.NewUsers(db)

	cfgPostgres := postgres.ConfigPostgres{
		//"localhost",
		"fullstack-postgres",
		"5432",
		"postgres",
		"password",
		"postgres",
		"disable",
	}
	pdb, err := postgres.NewPostgresDB(cfgPostgres)
	if err != nil {
		loggerService.Fatalf("!!!!!!DB is not connected: %s", err.Error())
	}
	postgresScheduleRepo := postgres.NewScheduleRepo(pdb)
	userRepo := postgres.NewUserRepo(pdb)

	r := repository.NewRepository(postgresScheduleRepo, userRepo)

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
		loggerService.Fatalf("Failed to stop HTTP server: %s", err)
	}

	// TODO: DISCONNECT FROM DB
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	//viper.
	return viper.ReadInConfig()
}
