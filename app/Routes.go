package app

import (
	"main/controller"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

func defineRoutes(e *echo.Echo) {

	e.GET("/hello", controller.HelloFunc)
	e.GET("/employee", controller.AllEmployees)
	e.GET("/department", controller.AllDepartments)

	e.GET("/api/employee/all", controller.GetAllEmployees)
	e.GET("/api/employee", controller.GetEmployeeById)
	e.GET("/api/employee/email", controller.GetEmployeeByEmail)
	e.POST("/api/employee", controller.CreateEmployeeV2)
	e.PUT("/api/employee", controller.UpdateEmployee)
	e.DELETE("/api/employee", controller.DeleteEmployee)

	e.GET("/api/department/all", controller.GetDepartments)
	e.GET("/api/department", controller.GetDepartmentById)
	e.GET("/api/department/name", controller.GetDepartmentByName)
	e.POST("/api/department", controller.CreateDepartment)
	e.PATCH("/api/department", controller.UpdateDepartment)
	e.DELETE("/api/department", controller.DeleteDepartment)

	e.POST("/api/empdep", controller.AddEmployeeToDepartment)
	e.PATCH("/api/empdep", controller.UpdateEffectFromDate)
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
