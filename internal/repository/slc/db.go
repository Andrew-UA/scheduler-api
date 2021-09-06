package slc

import "scheduler/internal/model"

type DB struct {
	Schedules         []model.ScheduleEvent
	Users             []model.User
	ScheduleIncrement int
	UserIncrement     int
}
