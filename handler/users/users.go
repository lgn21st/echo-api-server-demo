package users

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/lgn21st/echo-api-server-demo/model"
	"github.com/lgn21st/echo-api-server-demo/service"
)

type response struct {
	Result string `json:"result"`
}

type tokenResponse struct {
	Result string `json:"result"`
	Token  string `json:"token"`
}

type updateRequest struct {
	Name string `json:"name" form:"name" query:"name"`
}

func Create(c echo.Context) error {
	u := &model.User{}
	c.Bind(u)

	err := service.CreateUser(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Result: err.Error()})
	} else {
		return c.JSON(http.StatusOK, response{Result: "ok"})
	}
}

func Auth(c echo.Context) error {
	u := &model.User{}
	c.Bind(u)

	err := service.AuthUser(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Result: err.Error()})
	}

	token, err := service.IssueToken(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Result: err.Error()})
	}

	return c.JSON(http.StatusOK, tokenResponse{Result: "ok", Token: token})
}

func Update(c echo.Context) error {
	r := &updateRequest{}
	c.Bind(r)

	email := userEmailFromToken(c)
	user, err := service.FindUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, response{Result: err.Error()})
	}

	err = service.UpdateName(user, r.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{Result: err.Error()})
	}

	return c.JSON(http.StatusOK, response{Result: "ok"})
}

func userEmailFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["email"].(string)
}
