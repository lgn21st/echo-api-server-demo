package service

import (
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lgn21st/echo-api-server-demo/db"
	"github.com/lgn21st/echo-api-server-demo/model"
)

var JWTTokenSecret = []byte("MyDarkSecret")

func CreateUser(u *model.User) error {
	if err := db.DB.Create(u).Error; err != nil {
		return &createError{user: u}
	}

	return nil
}

func FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}

	if err := db.DB.Where(&model.User{Email: email}).First(&user).Error; err != nil {
		return nil, &findError{Email: email}
	}

	return user, nil
}

func AuthUser(u *model.User) error {
	user, err := FindUserByEmail(u.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(u.Password))
	if err != nil {
		return &authError{user: u}
	}

	return nil
}

func IssueToken(u *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email

	return token.SignedString(JWTTokenSecret)
}

func UpdateName(u *model.User, name string) error {
	u.Name = name
	if err := db.DB.Save(u).Error; err != nil {
		return &updateError{user: u}
	}

	return nil
}
