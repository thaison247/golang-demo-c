package structs

type EmpDep struct {
	Id           int `json:"id" query:"id" db:"id"`
	EmployeeId   int `json:"employee_id" query:"employeeid" db:"employee_id"`
	DepartmentId int `json:"department_id" query:"departmentid" db:"department_id"`
	// EffectFrom   time.Time `json:"effect_from" query:"effectfrom" db:"effect_from"`
	// CreatedAt    time.Time `json:"created_at" query:"createdat" db:"created_at"`
}
