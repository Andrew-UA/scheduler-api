package postgres

import (
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"github.com/jmoiron/sqlx"
	"scheduler/internal/model"
	"strconv"
	"time"
)

const (
	param_interval_day   = "day"
	param_interval_week  = "week"
	param_interval_month = "month"
)

type ScheduleRepo struct {
	db *sqlx.DB
}

func NewScheduleRepo(db *sqlx.DB) *ScheduleRepo {
	return &ScheduleRepo{
		db: db,
	}
}

func (r *ScheduleRepo) List(params map[string]string) ([]model.ScheduleEvent, error) {
	whereStatements := make([]string, 0, 2)
	userId, isExists := params["user_id"]
	if isExists {
		userIdInt, err := strconv.Atoi(userId)
		if err != nil {
			return nil, err
		}
		whereStatements = append(whereStatements, fmt.Sprintf(" user_id = %d ", userIdInt))
	}

	interval, isExists := params["interval"]
	if isExists {
		from, to, err := getIntervalDates(interval)
		if err == nil {
			whereStatements = append(whereStatements, fmt.Sprintf(" start_at >= %d AND start_at <= %d ", from, to))
		}
	}

	query := "SELECT * FROM schedule_events"
	if len(whereStatements) != 0 {
		for key, value := range whereStatements {
			if key == 0 {
				query += " WHERE " + value
			} else {
				query += " AND " + value
			}
		}
	}
	query += " ORDER BY created_at"
	scheduleEvents := []model.ScheduleEvent{}
	err := r.db.Select(&scheduleEvents, query)
	if err != nil {
		return nil, err
	}

	return scheduleEvents, nil
}
func (r *ScheduleRepo) Show(ID int) (model.ScheduleEvent, error) {
	scheduleEvent := model.ScheduleEvent{}
	query := fmt.Sprintf("SELECT * FROM schedule_events WHERE id = %d LIMIT 1", ID)
	err := r.db.Get(&scheduleEvent, query)
	if err != nil {
		return model.ScheduleEvent{}, err
	}

	return scheduleEvent, nil
}

func (r *ScheduleRepo) Create(m model.ScheduleEvent) (model.ScheduleEvent, error) {
	query := "INSERT INTO schedule_events (user_id, name, time, start_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	row := r.db.QueryRow(query,
		m.UserID,
		m.Name,
		m.Time,
		m.StartAt,
		time.Now().Unix(),
		time.Now().Unix(),
	)

	schedule := model.ScheduleEvent{}
	err := row.Scan(
		&schedule.ID,
		&schedule.UserID,
		&schedule.Name,
		&schedule.Time,
		&schedule.StartAt,
		&schedule.CreatedAt,
		&schedule.UpdatedAt,
	)
	if err != nil {
		return model.ScheduleEvent{}, err
	}

	return schedule, nil
}
func (r *ScheduleRepo) Update(ID int, m model.ScheduleEvent) (model.ScheduleEvent, error) {

	mInDB, err := r.Show(ID)
	if err != nil {
		return model.ScheduleEvent{}, err
	}

	if mInDB.Name == m.Name && mInDB.StartAt == m.StartAt && mInDB.Time == m.Time {
		return mInDB, nil
	}

	m.UpdatedAt = time.Now().Unix()
	query := fmt.Sprintf("UPDATE schedule_events SET name = $1, time = $2, start_at = $3, updated_at = $4 WHERE id = %d", ID)
	_, err = r.db.Exec(query,
		m.Name,
		m.Time,
		m.StartAt,
		m.UpdatedAt,
	)
	if err != nil {
		return model.ScheduleEvent{}, err
	}
	m.ID = mInDB.ID
	m.UserID = mInDB.UserID
	m.CreatedAt = mInDB.CreatedAt

	return m, nil
}
func (r *ScheduleRepo) Delete(ID int) error {
	query := fmt.Sprintf("DELETE FROM schedule_events WHERE id = %d", ID)
	_, err := r.db.Exec(query)

	return err
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
		return 0, 0, errors.New("INVALID INTERVAL")
	}
}
