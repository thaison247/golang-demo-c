package app

import (
	"main/controller"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func defineRoutes(e *echo.Echo) {
	e.GET("/hello", controller.HelloFunc)
	e.GET("/api/employee/all", controller.GetAllEmployee)
	e.GET("/api/employee", controller.GetEmployeeById)
	e.POST("/api/employee", controller.CreateEmployee)
	e.PATCH("api/employee", controller.UpdateEmployee)
	e.DELETE("api/employee", controller.DeleteEmployee)
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
