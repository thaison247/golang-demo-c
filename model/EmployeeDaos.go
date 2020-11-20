package model

import (
	database "g.ghn.vn/scte-common/godal"
)

const SQL_CUSTOM_GET_BY_ID = "" +
	"SELECT * " +
	"FROM employees " +
	"WHERE employee_id = $1 "

func GetEmployeeById(dbType database.IDatabase, id int) (interface{}, error) {
	params := []interface{}{id}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_BY_ID, params)
}
