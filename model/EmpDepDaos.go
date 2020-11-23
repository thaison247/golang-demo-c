package model

import (
	database "g.ghn.vn/scte-common/godal"
)

var (
	EMP_DEP string = "emp_dep"
)

const SQL_CUSTOM_GET_EMPLOYEE_DEPARTMENT = "" + "SELECT * FROM get_employee_departmentid($1)"

func GetEmployeeDepartmentId(dbType database.IDatabase, employeeId int) ([]map[string]interface{}, error) {
	params := []interface{}{employeeId}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEE_DEPARTMENT, params)
}

func AddEmployeeToDeparmentWithMap(dbType database.IDatabase, mapData map[string]interface{}) (interface{}, error) {
	return dbType.Create(EMP_DEP, mapData)
}
