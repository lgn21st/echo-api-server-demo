package service

import (
	"fmt"

	"github.com/lgn21st/echo-api-server-demo/db"
	"github.com/lgn21st/echo-api-server-demo/models"
)

type userCreateError struct {
	user *models.User
}

func (e *userCreateError) Error() string {
	return fmt.Sprintf("Create user failed for User{name: %v, email: %v}", e.user.Name, e.user.Email)
}

func CreateUser(u *models.User) error {
	db.DB.Create(u)

	if db.DB.NewRecord(u) {
		return &userCreateError{user: u}
	} else {
		return nil
	}
}
