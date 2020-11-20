package structs

type Department struct {
	Id     int    `json:"-" query:"departmentid" db:"department_id"`
	Code   string `json:"department_code" query:"departmentcode" db:"department_code"`
	Name   string `json:"department_name" query:"departmentname" db:"department_name"`
	Active bool   `json:"active" query:"active" db:"active"`
}
