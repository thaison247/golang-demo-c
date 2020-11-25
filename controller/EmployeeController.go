package controller

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/structs"
	"main/utils"
	"net/http"
	"strconv"

	database "g.ghn.vn/scte-common/godal"
	"github.com/labstack/echo/v4"
)

var (
	EMPLOYEES string = "employees"
)

func CreateEmployee(c echo.Context) error {
	var err error

	dataReq := new(structs.Employee)
	if err = c.Bind(dataReq); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	res, err := model.Create(dbType, EMPLOYEES, dataReq)

	fmt.Println(res)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, "Success")
}

func CreateEmployeeV2(c echo.Context) error {
	var err error

	dataReq := new(structs.Employee)
	if err = c.Bind(dataReq); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	jsonData, err := json.Marshal(dataReq)
	var newData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &newData)

	delete(newData, "employee_id")

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	res, err := model.CreateWithMap(dbType, newData)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, res)
}

func GetAllEmployees(c echo.Context) error {
	var limit int
	var offset int
	var err error

	if limit, err = strconv.Atoi(c.QueryParam("limit")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	if offset, err = strconv.Atoi(c.QueryParam("offset")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.GetEmployeesWithDepartment(dbType, limit, offset)

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

	res, err := model.Update(dbType, EMPLOYEES, newData, map[string]interface{}{"employee_id": employeeId})

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
	_, err = model.Delete(dbType, EMPLOYEES, map[string]interface{}{"employee_id": employeeId})

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, "success")
}
