package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lgn21st/echo-api-server-demo/db"
	"github.com/lgn21st/echo-api-server-demo/models"
)

type response struct {
	Result string `json:"result"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/users", func(c echo.Context) error {
		u := &models.User{}
		c.Bind(u)
		db.DB.Create(&u)

		return c.JSON(http.StatusOK, response{Result: "ok"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
