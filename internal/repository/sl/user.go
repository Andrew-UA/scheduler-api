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

func (u *Users) FindByLogin(login string) (model.User, error) {
	for _, user := range u.DB.Users {
		if login == user.Login {
			return user, nil
		}
	}

	return model.User{}, errors.New("NOT FOUND")
}

func (u *Users) FindByID(ID int) (model.User, error) {
	for _, user := range u.DB.Users {
		if ID == user.ID {
			return user, nil
		}
	}

	return model.User{}, errors.New("NOT FOUND")
}

func (u *Users) Update(userID int, data model.User) (model.User, error) {
	for key, user := range u.DB.Users {
		if user.ID == userID {
			u.DB.Users[key].Timezone = data.Timezone
		}

		return u.DB.Users[key], nil
	}

	return model.User{}, errors.New("NOT FOUND")
}
