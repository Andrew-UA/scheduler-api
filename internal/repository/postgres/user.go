package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"scheduler/internal/model"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindByLogin(login string) (model.User, error) {
	user := model.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE login = '%s' LIMIT 1", login)
	err := r.db.Get(&user, query)

	return user, err
}

func (r *UserRepo) FindByID(ID int) (model.User, error) {
	user := model.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d LIMIT 1", ID)
	err := r.db.Get(&user, query)

	return user, err
}

func (r *UserRepo) Update(userID int, user model.User) (model.User, error) {

	userInDB, err := r.FindByID(userID)
	if err != nil {
		return user, err
	}

	user.ID = userInDB.ID
	user.Login = userInDB.Login
	if userInDB.Timezone == user.Timezone {
		return user, nil
	}

	query := fmt.Sprintf("UPDATE users SET timezone = $1 WHERE id = %d", userID)
	_, err = r.db.Exec(query, user.Timezone)

	return user, err
}
