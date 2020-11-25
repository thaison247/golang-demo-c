package structs

type Department struct {
	Id     int    `json:"department_id" query:"departmentid" db:"department_id"`
	Code   string `json:"department_code" query:"departmentcode" db:"department_code"`
	Name   string `json:"department_name" query:"departmentname" db:"department_name"`
	Active bool   `json:"active" query:"active" db:"active"`
	// CreatedAt time.Time `json:"created_at" query:"createdat" db:"created_at"`
	// UpdatedAt time.Time `json:"updated_at" query:"updatedat" db:"updated_at"`
}
