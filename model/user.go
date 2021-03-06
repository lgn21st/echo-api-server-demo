package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// The User Model
type User struct {
	gorm.Model
	Name              string `json:"name" form:"name" query:"name"`
	Email             string `json:"email" form:"email" query:"email"`
	Password          string `json:"password" form:"password" query:"password" sql:"-"`
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

// TODO: Move encryptPassword() to service instead of Gorm callback?
func (u *User) BeforeCreate() error {
	return u.encryptPassword()
}
