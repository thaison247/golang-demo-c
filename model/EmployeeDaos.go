package model

import (
	database "g.ghn.vn/scte-common/godal"
)

var (
	EMPLOYEES = "employees"
)

const SQL_CUSTOM_GET_EMPLOYEE_BY_ID = "" + "SELECT * FROM get_one_employee_with_department_v2($1)"

const SQL_CUSTOM_GET_EMPLOYEES_WITH_DEPARTMENT = "" + "SELECT * " + "FROM get_employees_with_department_v2($1, $2)"

const SQL_CUSTOM_GET_EMPLOYEE_BY_EMAIL = "" + "SELECT * FROM employees WHERE email = ($1)"

func GetEmployeeById(dbType database.IDatabase, id int) (interface{}, error) {
	params := []interface{}{id}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEE_BY_ID, params)
}

func GetEmployeesWithDepartment(dbType database.IDatabase, limit int, offset int) ([]map[string]interface{}, error) {
	params := []interface{}{limit, offset}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEES_WITH_DEPARTMENT, params)
}

func CreateWithMap(dbType database.IDatabase, mapData map[string]interface{}) (interface{}, error) {
	return dbType.Create(EMPLOYEES, mapData)
}

func GetEmployeeByEmail(dbType database.IDatabase, email string) (interface{}, error) {
	params := []interface{}{email}
	return dbType.ExecuteSelectToMap(SQL_CUSTOM_GET_EMPLOYEE_BY_EMAIL, params)
}
