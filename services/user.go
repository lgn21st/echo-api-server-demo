package services

import (
	"fmt"

	"github.com/lgn21st/echo-api-server-demo/db"
	"github.com/lgn21st/echo-api-server-demo/models"
)

type userCreateError struct {
	user *models.User
}

func (e *userCreateError) Error() string {
	return fmt.Sprintf("Create User error.")
}

func CreateUser(u *models.User) error {
	db.DB.Create(u)

	if db.DB.NewRecord(u) {
		return &userCreateError{user: u}
	} else {
		return nil
	}
}
