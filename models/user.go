package models

import "github.com/jinzhu/gorm"

// The User Model
type User struct {
	gorm.Model
	Email             string `json:"email" form:"email" query:"email"`
	EncryptedPassword string `json:"password" form:"password" query:"password"`
	Name              string `json:"name" form:"name" query:"name"`
}
