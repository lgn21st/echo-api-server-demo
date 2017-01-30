package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// The User Model
type User struct {
	gorm.Model
	Name              string `json:"name" form:"name" query:"name"`
	Email             string `json:"email" form:"email" query:"email"`
	Password          string `sql:"-"`
	EncryptedPassword string
}

func (u *User) encryptPassword() error {
	pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.EncryptedPassword = string(pw)
	return nil
}

func (u *User) BeforeCreate() error {
	return u.encryptPassword()
}
