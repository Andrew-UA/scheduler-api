package model

import (
	"encoding/json"
	"time"
)

const (
	layout = "01/02/06 15:04:05"
)

type ScheduleEvent struct {
	ID        int    `json:"id" db:"id"`
	UserID    int    `json:"user_id" db:"user_id"`
	Name      string `json:"name" db:"name"`
	Time      int    `json:"time" db:"time"`
	StartAt   int64  `json:"start_at" db:"start_at"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
}

type ScheduleEventJson struct {
	ScheduleEvent
	StartAt   string `json:"start_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *ScheduleEvent) Marshal(timezone string) ([]byte, error) {
	scheduleJson, err := s.ToScheduleEventJson(timezone)
	if err != nil {
		return nil, err
	}
	jsonString, err := json.Marshal(scheduleJson)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func (s *ScheduleEvent) ToScheduleEventJson(timezone string) (*ScheduleEventJson, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	return &ScheduleEventJson{
		ScheduleEvent: *s,
		StartAt:       time.Unix(s.StartAt, 0).In(location).Format(layout),
		CreatedAt:     time.Unix(s.CreatedAt, 0).In(location).Format(layout),
		UpdatedAt:     time.Unix(s.UpdatedAt, 0).In(location).Format(layout),
	}, nil
}

func (s *ScheduleEvent) FromScheduleEventJson(eventJson ScheduleEventJson, timezone string) error {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}
	s.Time = eventJson.Time
	s.Name = eventJson.Name
	t, err := time.ParseInLocation(layout, eventJson.StartAt, location)
	if err != nil {
		return err
	}
	s.StartAt = t.Unix()

	return nil
}
