package auth

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IPasswordManager interface {
	HashAndSalt(pwd string) (string, error)
	CheckPassword(hashedPwd, pwd string) bool
}

type PasswordManager struct {
	cost int
}

func NewPasswordManager() *PasswordManager {
	return &PasswordManager{
		cost: bcrypt.MinCost,
	}
}

func (pm *PasswordManager) HashAndSalt(pwd string) (string, error) {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), pm.cost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func (pm *PasswordManager) CheckPassword(hashedPwd, pwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
