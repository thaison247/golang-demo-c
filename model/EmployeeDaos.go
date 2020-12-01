package model

import (
	database "g.ghn.vn/scte-common/godal"
)

var (
	EMPLOYEES = "employees"
)

const SQL_CUSTOM_GET_EMPLOYEE_BY_ID = "" + "SELECT * FROM get_one_employee_with_department_v3($1)"

const SQL_CUSTOM_GET_EMPLOYEES_WITH_DEPARTMENT = "" + "SELECT * " + "FROM get_employees_with_department_v3($1, $2)"

const SQL_CUSTOM_GET_EMPLOYEE_BY_EMAIL = "" + "SELECT * FROM employees WHERE email = ($1)"

const SQL_CUSTOM_ADD_EMPLOYEE = "CALL add_employee($1::text, $2::text, $3::text, $4, $5::text, $6::text, $7, $8::date)"

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

func AddEmployee(dbType database.IDatabase, employee map[string]interface{}) (interface{}, error) {
	fullName := employee["full_name"]
	phoneNumber := employee["phone_number"]
	email := employee["email"]
	gender := employee["gender"]
	jobTitle := employee["job_title"]
	address := employee["address"]
	departmentId := employee["department_id"]
	effectFrom := employee["effect_from"]
	params := []interface{}{fullName, phoneNumber, email, gender, jobTitle, address, departmentId, effectFrom}
	return dbType.Execute(SQL_CUSTOM_ADD_EMPLOYEE, params)
}
