package model

import (
	database "g.ghn.vn/scte-common/godal"
)

var (
	EMPLOYEES = "employees"
)

const SQL_CUSTOM_GET_EMPLOYEE_BY_ID = "" + "SELECT * FROM get_employee_by_id($1)"

const SQL_CUSTOM_GET_EMPLOYEE_WITH_DEPARTMENTID = "" + "SELECT * " + "FROM get_employees_with_departmentid($1, $2)"

const SQL_CUSTOM_GET_EMPLOYEES_WITH_DEPARTMENT = "" + "SELECT * " + "FROM get_employees_with_department($1, $2)"

func GetEmployeeById(dbType database.IDatabase, id int) (interface{}, error) {
	params := []interface{}{id}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEE_BY_ID, params)
}

func GetEmployeesWithDepartmentId(dbType database.IDatabase, limit int, offset int) ([]map[string]interface{}, error) {
	params := []interface{}{limit, offset}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEE_WITH_DEPARTMENTID, params)
}

func GetEmployeesWithDepartment(dbType database.IDatabase, limit int, offset int) ([]map[string]interface{}, error) {
	params := []interface{}{limit, offset}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEES_WITH_DEPARTMENT, params)
}

func CreateWithMap(dbType database.IDatabase, mapData map[string]interface{}) (interface{}, error) {
	return dbType.Create(EMPLOYEES, mapData)
}
