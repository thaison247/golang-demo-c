package app

import (
	"main/controller"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func defineRoutes(e *echo.Echo) {
	e.GET("/hello", controller.HelloFunc)
	e.POST("/api/user", controller.CreateUser)
	e.GET("/api/user", controller.GetUser)
}

func Routes(e *echo.Echo) {
	// Define all routes of API here
	defineRoutes(e)

	// Print all the routes of API on
	allRoutes := e.Routes()
	for i := 0; i < len(allRoutes); i++ {
		route := allRoutes[i]
		log.Infof("%s %s", route.Method, route.Path)
	}
}
