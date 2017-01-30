package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/lgn21st/echo-api-server-demo/db"
	"github.com/lgn21st/echo-api-server-demo/model"
)

var jwtTokenSecret = []byte("MyDarkSecret")

type createError struct {
	user *model.User
}

func (e *createError) Error() string {
	return fmt.Sprintf("Create user failed for User{name: %v, email: %v}", e.user.Name, e.user.Email)
}

type findError struct {
	user *model.User
}

func (e *findError) Error() string {
	return fmt.Sprintf("User not found for User{email: %v}", e.user.Email)
}

type authError struct {
	user *model.User
}

func (e *authError) Error() string {
	return fmt.Sprintf("User auth failed for User{email: %v, password: %v}", e.user.Email, e.user.Password)
}

func CreateUser(u *model.User) error {
	db.DB.Create(u)

	if db.DB.NewRecord(u) {
		return &createError{user: u}
	} else {
		return nil
	}
}

func FindUser(u *model.User) (*model.User, error) {
	user := &model.User{}

	if db.DB.Where(&model.User{Email: u.Email}).First(&user).Error != nil {
		return nil, &findError{user: u}
	}

	return user, nil
}

func AuthUser(u *model.User) error {
	user, err := FindUser(u)
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
	claims["name"] = u.Name

	return token.SignedString(jwtTokenSecret)
}
