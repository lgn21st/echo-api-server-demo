package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type response struct {
	Result string `json:"result"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{Result: "ok"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
