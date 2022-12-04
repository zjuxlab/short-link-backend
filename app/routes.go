package app

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ping(c echo.Context) error {
	return c.String(200, "pong!")
}

func addRoutes() {
	api := e.Group("api")
	api.GET("/doc/*", echoSwagger.WrapHandler)
	api.GET("/ping", ping)
}
