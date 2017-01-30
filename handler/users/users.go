package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lgn21st/echo-api-server-demo/model"
	"github.com/lgn21st/echo-api-server-demo/service"
)

type Response struct {
	Result string `json:"result"`
}

type TokenResponse struct {
	Result string `json:"result"`
	Token  string `json:"token"`
}

func Create(c echo.Context) error {
	u := &model.User{}
	c.Bind(u)

	err := service.CreateUser(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Result: err.Error()})
	} else {
		return c.JSON(http.StatusOK, Response{Result: "ok"})
	}
}

func Auth(c echo.Context) error {
	u := &model.User{}
	c.Bind(u)

	err := service.AuthUser(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Result: err.Error()})
	}

	token, err := service.IssueToken(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Result: err.Error()})
	}

	return c.JSON(http.StatusOK, TokenResponse{Result: "ok", Token: token})
}
