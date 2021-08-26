package auth

import (
	"errors"
	"scheduler/internal/repository"
	"scheduler/pkg/auth"
	"strconv"
	"time"
)

type Authenticator interface {
	SignIn(login, password string) (string, error)
}

type JwtService struct {
	Repo            repository.IUser
	passwordManager auth.IPasswordManager
	tokeManager     auth.TokenManager
}

func NewJwtService(repo repository.IUser) *JwtService {
	return &JwtService{
		Repo:            repo,
		passwordManager: auth.NewPasswordManager(),
		tokeManager:     auth.NewTokenManager(""),
	}
}

func (s *JwtService) SignIn(login, password string) (string, error) {
	user, rErr := s.Repo.FindByLogin(login)
	if rErr != nil {
		return "", rErr
	}

	if !s.passwordManager.CheckPassword(user.Password, password) {
		return "", errors.New("INCORRECT LOGIN OR PASSWORD")
	}

	token, tErr := s.tokeManager.NewJWT(strconv.Itoa(user.ID), time.Minute*5)
	if tErr != nil {
		return "", tErr
	}

	return token, nil
}
