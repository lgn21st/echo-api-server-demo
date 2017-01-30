package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/lgn21st/echo-api-server-demo/handler"
	userHandler "github.com/lgn21st/echo-api-server-demo/handler/users"
)

func main() {
	api := echo.New()

	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())

	api.GET("/ping", handler.Ping)
	api.POST("/users", userHandler.Create)

	api.Logger.Fatal(api.Start(":8080"))
}
