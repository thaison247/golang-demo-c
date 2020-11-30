package model

import (
	database "g.ghn.vn/scte-common/godal"
)

const DEPARTMENTS = "departments"

const SQL_CUSTOM_GET_DEPARTMENT_BY_ID = "" +
	"SELECT * " +
	"FROM departments " +
	"WHERE department_id = $1 "

const SQL_CUSTOM_GET_DEPARTMENT_BY_NAME = "SELECT * FROM departments WHERE department_name = $1"

func GetDepartmentById(dbType database.IDatabase, id int) (interface{}, error) {
	params := []interface{}{id}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_DEPARTMENT_BY_ID, params)
}

func GetDepartmentByName(dbType database.IDatabase, departmentName string) (interface{}, error) {
	params := []interface{}{departmentName}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_DEPARTMENT_BY_NAME, params)
}

func CreateDepartmentWithMap(dbType database.IDatabase, mapData map[string]interface{}) (interface{}, error) {
	return dbType.Create(DEPARTMENTS, mapData)
}
