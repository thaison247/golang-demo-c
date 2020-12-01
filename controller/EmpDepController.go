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
	fmt.Println(newData)
	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)

	// // get latest effect day
	// result, err := model.GetLatestEffectDay(dbType, int(newData["employee_id"].(float64)))
	// fmt.Printf("res: %+v", result)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ApiResult(c, http.StatusBadRequest, err)
	// }

	// // t := time.Now()

	// latestDay := result.(map[string]interface{})["effect_from"].(time.Time)
	// inputDay := newData["effect_from"].(time.Time)

	// diff := latestDay.Sub(inputDay)

	// if diff > 0 {
	// 	fmt.Println(diff)
	// 	return ApiResult(c, http.StatusBadRequest, "invalid effect day")
	// }

	delete(newData, "id")
	_, err = model.AddEmployeeToDeparmentWithMap(dbType, newData)

	if err != nil {
		fmt.Println("res error: ", err)
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

func GetLatestEffectDayByEmpId(c echo.Context) error {
	var employeeId int
	var err error

	if employeeId, err = strconv.Atoi(c.QueryParam("employeeid")); err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	dbType := utils.Global[utils.POSTGRES_ENTITY].(database.Postgres)
	rs, err := model.GetLatestEffectDay(dbType, employeeId)

	if err != nil {
		return ApiResult(c, http.StatusBadRequest, err)
	}

	return ApiResult(c, http.StatusOK, rs)
}
