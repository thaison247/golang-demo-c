package controller

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/structs"
	"main/utils"
	"net/http"

	database "g.ghn.vn/scte-common/godal"
	"github.com/labstack/echo/v4"
)

var (
	EMP_DEP string = "emp_dep"
)

func AddEmployeeToDepartment(c echo.Context) error {
	var err error

	// get data from request, bind to EmpDep struct
	dataReq := new(structs.EmpDep)
	if err = c.Bind(dataReq); err != nil {
		fmt.Println(err)
		return ApiResult(c, http.StatusBadRequest, err)
	}

	//parse request data to map[string]interface{}
	jsonData, err := json.Marshal(dataReq)
	var newData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &newData)

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)

	delete(newData, "id")
	_, err = model.AddEmployeeToDeparmentWithMap(dbType, newData)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, "success")
}

func UpdateEffectFromDate(c echo.Context) error {
	var err error

	// get data from request, bind to EmpDep struct
	dataReq := new(structs.EmpDep)
	if err = c.Bind(dataReq); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	//parse request data to map[string]interface{}
	jsonData, err := json.Marshal(dataReq)
	var newData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &newData)

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	_, err = model.Update(dbType, EMP_DEP, newData, map[string]interface{}{"id": dataReq.Id})
	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, "success")
}
