package model

import (
	database "g.ghn.vn/scte-common/godal"
)

const SQL_CUSTOM_GET_DEPARTMENT_BY_ID = "" +
	"SELECT * " +
	"FROM departments " +
	"WHERE department_id = $1 "

func GetDepartmentById(dbType database.IDatabase, id int) (interface{}, error) {
	params := []interface{}{id}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_DEPARTMENT_BY_ID, params)
}
