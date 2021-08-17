package sl

import "scheduler/internal/model"

type DB struct {
	Schedules []model.ScheduleEvent
	Users []string
	ScheduleIncrement int
}
