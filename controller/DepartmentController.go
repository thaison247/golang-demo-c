package controller

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/structs"
	"main/utils"
	"net/http"
	"strconv"
	"time"

	database "g.ghn.vn/scte-common/godal"
	"github.com/labstack/echo/v4"
)

var (
	DEPARTMENTS string = "departments"
)

func CreateDepartment(c echo.Context) error {
	var err error

	dataReq := new(structs.Department)
	if err = c.Bind(dataReq); err != nil {
		fmt.Println(err)
		return ApiResult(c, http.StatusBadRequest, err)
	}

	jsonData, err := json.Marshal(dataReq)
	var newData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &newData)

	delete(newData, "department_id")
	fmt.Println(newData)

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err = model.CreateDepartmentWithMap(dbType, newData)
	if err != nil {
		fmt.Println(err)
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, "Success")
}

func GetDepartments(c echo.Context) error {
	var limit int
	var offset int
	var err error

	if limit, err = strconv.Atoi(c.QueryParam("limit")); err != nil {
		ApiResult(c, http.StatusBadRequest, err)
	}

	if offset, err = strconv.Atoi(c.QueryParam("offset")); err != nil {
		ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.Get(dbType, DEPARTMENTS, limit, offset)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}

func GetDepartmentById(c echo.Context) error {
	var departmentId int
	var err error

	if departmentId, err = strconv.Atoi(c.QueryParam("departmentid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.GetDepartmentById(dbType, departmentId)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}

func GetDepartmentByName(c echo.Context) error {

	departmentName := c.QueryParam("departmentname")

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.GetDepartmentByName(dbType, departmentName)

	if err != nil {
		fmt.Println(err)
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}

func UpdateDepartment(c echo.Context) error {
	var departmentId int
	var err error

	if departmentId, err = strconv.Atoi(c.QueryParam("departmentid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dataReq := new(structs.Department)
	if err = c.Bind(dataReq); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	jsonData, err := json.Marshal(dataReq)
	var newData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &newData)
	currentTime := time.Now()
	newData["updated_at"] = string(currentTime.Format(time.RFC3339))

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)

	res, err := model.Update(dbType, DEPARTMENTS, newData, map[string]interface{}{"department_id": departmentId})

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, res)
}

func DeleteDepartment(c echo.Context) error {
	var err error
	var departmentId int

	if departmentId, err = strconv.Atoi(c.QueryParam("departmentid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	res, err := model.Delete(dbType, DEPARTMENTS, map[string]interface{}{"department_id": departmentId})

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, res)
}
