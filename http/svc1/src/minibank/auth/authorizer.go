package minibank

import (
	models "minibank/models"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
}

func HashAccount(a *models.Account) (models.Account, error) {
	var err error
	a.Secret, err = hashPassword(a.Secret)
	return *a, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
