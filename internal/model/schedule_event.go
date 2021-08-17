package model

import (
	"encoding/json"
	"time"
)

const (
	layout = "01/02/06 15:04:05"
)

type ScheduleEvent struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Time      int    `json:"time"`
	StartAt   int64  `json:"start_at"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type ScheduleEventJson struct {
	ScheduleEvent
	StartAt   string `json:"start_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (s *ScheduleEvent) Marshal() ([]byte, error) {
	scheduleJson := s.ToScheduleEventJson()
	jsonString, err := json.Marshal(scheduleJson)
	if err != nil {
		return nil, err
	}

	return jsonString, nil
}

func (s *ScheduleEvent) ToScheduleEventJson() *ScheduleEventJson {
	return &ScheduleEventJson{
		ScheduleEvent: *s,
		StartAt:   time.Unix(s.StartAt, 0).Format(layout),
		CreatedAt: time.Unix(s.CreatedAt, 0).Format(layout),
		UpdatedAt: time.Unix(s.UpdatedAt, 0).Format(layout),
	}
}

func (s *ScheduleEvent) FromScheduleEventJson(eventJson ScheduleEventJson) error {
	s.Time = eventJson.Time
	s.Name = eventJson.Name
	t, err := time.Parse(layout, eventJson.StartAt)
	if err != nil {
		return err
	}
	s.StartAt = t.Unix()

	return nil
}