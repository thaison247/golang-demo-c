package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiResult(c echo.Context, status int, data interface{}) error {
	mapData := make(map[string]interface{})
	mapData["status"] = status
	mapData["data"] = data

	return c.JSON(status, mapData)
}

func HelloFunc(c echo.Context) error {
	welcomeMess := fmt.Sprintf("Welcome to HRTech Go Core")
	return ApiResult(c, http.StatusOK, welcomeMess)
}
