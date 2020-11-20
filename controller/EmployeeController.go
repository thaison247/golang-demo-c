package controller

import (
	"encoding/json"
	"main/model"
	"main/structs"
	"main/utils"
	"net/http"
	"strconv"

	database "g.ghn.vn/scte-common/godal"
	"github.com/labstack/echo/v4"
)

var (
// tableName string = "employees"
)

func CreateEmployee(c echo.Context) error {
	var err error

	dataReq := new(structs.Employee)
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

func GetAllEmployees(c echo.Context) error {
	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.Get(dbType, tableName, 10, 0)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}

func GetEmployeeById(c echo.Context) error {
	var employeeId int
	var err error

	if employeeId, err = strconv.Atoi(c.QueryParam("employeeid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.GetEmployeeById(dbType, employeeId)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}

func UpdateEmployee(c echo.Context) error {
	var employeeId int
	var err error

	if employeeId, err = strconv.Atoi(c.QueryParam("employeeid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dataReq := new(structs.Employee)
	if err = c.Bind(dataReq); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	jsonData, err := json.Marshal(dataReq)
	var newData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &newData)

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)

	res, err := model.Update(dbType, tableName, newData, map[string]interface{}{"employee_id": employeeId})

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, res)
}

func DeleteEmployee(c echo.Context) error {
	var err error
	var employeeId int

	if employeeId, err = strconv.Atoi(c.QueryParam("employeeid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	res, err := model.Delete(dbType, tableName, map[string]interface{}{"employee_id": employeeId})

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, res)
}
