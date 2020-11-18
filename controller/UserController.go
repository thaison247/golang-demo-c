package controller

import (
	"main/model"
	"main/structs"
	"main/utils"
	"net/http"
	"strconv"

	database "g.ghn.vn/scte-common/godal"
	"github.com/labstack/echo/v4"
)

var (
	tableName string = "users"
)

func CreateUser(c echo.Context) error {
	var err error

	dataReq := new(structs.User)
	if err = c.Bind(dataReq); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err = model.Create(dbType, tableName, dataReq)
	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, "Success")
}

func GetUser(c echo.Context) error {
	var lenId int
	var err error

	if lenId, err = strconv.Atoi(c.QueryParam("len")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.CustomGet(dbType, lenId)
	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}
