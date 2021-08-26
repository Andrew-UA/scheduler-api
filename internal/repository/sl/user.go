package sl

import (
	"errors"
	"scheduler/internal/model"
)

type Users struct {
	DB *DB
}

func NewUsers(db *DB) *Users {
	return &Users{
		DB: db,
	}
}

func (u *Users)FindByLogin(login string)  (model.User, error) {
	for _, user := range u.DB.Users {
		if login == user.Login {
			return user, nil
		}
	}

	return model.User{}, errors.New("NOT FOUND")
}