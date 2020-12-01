package model

import (
	database "g.ghn.vn/scte-common/godal"
)

var (
	EMP_DEP string = "emp_dep"
)

const SQL_CUSTOM_GET_LATEST_EFFECT_DAY = "SELECT max(effect_from) as effect_from FROM emp_dep WHERE employee_id = $1"

func AddEmployeeToDeparmentWithMap(dbType database.IDatabase, mapData map[string]interface{}) (interface{}, error) {
	return dbType.Create(EMP_DEP, mapData)
}

func GetLatestEffectDay(dbType database.IDatabase, employeeId int) (interface{}, error) {
	params := []interface{}{employeeId}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_LATEST_EFFECT_DAY, params)
}
