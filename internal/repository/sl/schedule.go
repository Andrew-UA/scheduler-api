package sl

import (
	"errors"
	"github.com/jinzhu/now"
	"scheduler/internal/model"
	"time"
)

const (
	param_interval_day = "day"
	param_interval_week = "week"
	param_interval_month = "month"
)

var param_intervals = [3]string {
	param_interval_day,
	param_interval_week,
	param_interval_month,
}

type Schedules struct {
	DB *DB
}

func NewSchedules(db *DB) *Schedules {
	return &Schedules{
		DB: db,
	}
}

func (s *Schedules) List(params map[string]string) ([]model.ScheduleEvent, error) {
	schedules := s.DB.Schedules
	interval, isExists := params["interval"]
	if isExists {
		from, to, err := getIntervalDates(interval)
		if err == nil {
			newSchedules := make([]model.ScheduleEvent, 0, len(schedules))
			for _, schedule := range schedules {
				if schedule.StartAt >= from && schedule.StartAt <= to {
					newSchedules = append(newSchedules, schedule)
				}
			}
			schedules = newSchedules
		}
	}

	return schedules, nil
}
func (s *Schedules) Show(ID int) (model.ScheduleEvent, error)  {
	for _, scheduleEvent := range s.DB.Schedules {
		if scheduleEvent.ID == ID {
			return scheduleEvent, nil
		}
	}

	return model.ScheduleEvent{}, errors.New("NOT FOUND")
}

func (s *Schedules) Create(m model.ScheduleEvent) (model.ScheduleEvent, error) {
	s.DB.ScheduleIncrement++
	m.ID = s.DB.ScheduleIncrement
	m.CreatedAt = time.Now().Unix()
	m.UpdatedAt = time.Now().Unix()
	s.DB.Schedules = append(s.DB.Schedules, m)
	return m, nil
}
func (s *Schedules) Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error) {
	for key, scheduleEvent := range s.DB.Schedules {
		if scheduleEvent.ID == ID {
			beforeUpdate := scheduleEvent
			scheduleEvent.Name = m.Name
			scheduleEvent.Time = m.Time
			scheduleEvent.StartAt = m.StartAt
			if beforeUpdate != scheduleEvent {
				scheduleEvent.UpdatedAt = time.Now().Unix()
				s.DB.Schedules[key] = scheduleEvent
			}

			return scheduleEvent, nil
		}
	}

	return model.ScheduleEvent{}, errors.New("NOT FOUND")
}
func (s *Schedules) Delete(ID int) error {
	for key, scheduleEvent := range s.DB.Schedules {
		if scheduleEvent.ID == ID {
			l := len(s.DB.Schedules)
			s.DB.Schedules[key] = s.DB.Schedules[l-1]
			s.DB.Schedules = s.DB.Schedules[:l-1]

			return nil
		}
	}

	return errors.New("NOT FOUND")
}

func getIntervalDates(interval string) (from, to int64, err error) {
	switch interval {
	case param_interval_day:
		return now.BeginningOfDay().Unix(), now.EndOfDay().Unix(), nil
	case param_interval_week:
		return now.BeginningOfWeek().Unix(), now.EndOfWeek().Unix(), nil
	case param_interval_month:
		return now.BeginningOfMonth().Unix(), now.EndOfMonth().Unix(), nil
	default:
		return 0,0, errors.New("INVALID INTERVAL")
	}
}
