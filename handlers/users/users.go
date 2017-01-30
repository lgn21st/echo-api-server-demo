package users

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lgn21st/echo-api-server-demo/db"
	"github.com/lgn21st/echo-api-server-demo/models"
)

type createSuccessResponse struct {
	Result string `json:"result"`
}

func Create(c echo.Context) error {
	u := &models.User{}
	c.Bind(u)
	db.DB.Create(&u)

	return c.JSON(http.StatusOK, createSuccessResponse{Result: "ok"})
}
