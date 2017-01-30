package service

import (
	"fmt"

	"github.com/lgn21st/echo-api-server-demo/model"
)

type createError struct {
	user *model.User
}

func (e *createError) Error() string {
	return fmt.Sprintf("Create user failed for User{name: %v, email: %v}", e.user.Name, e.user.Email)
}

type findError struct {
	Email string
}

func (e *findError) Error() string {
	return fmt.Sprintf("User not found for User{email: %v}", e.Email)
}

type authError struct {
	user *model.User
}

func (e *authError) Error() string {
	return fmt.Sprintf("User auth failed for User{email: %v, password: %v}", e.user.Email, e.user.Password)
}

type updateError struct {
	user *model.User
}

func (e *updateError) Error() string {
	return fmt.Sprintf("Update failed for User{email: %v, name: %v}", e.user.Email, e.user.Name)
}
