package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lgn21st/echo-api-server-demo/models"
	"github.com/lgn21st/echo-api-server-demo/services"
)

type Response struct {
	Result string `json:"result"`
}

func Create(c echo.Context) error {
	u := &models.User{}
	c.Bind(u)

	err := services.CreateUser(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{Result: err.Error()})
	} else {
		return c.JSON(http.StatusOK, Response{Result: "ok"})
	}
}
