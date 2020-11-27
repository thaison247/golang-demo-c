package structs

type Employee struct {
	Id       int    `json:"employee_id" query:"employeeid" db:"employee_id"`
	Name     string `json:"full_name" query:"username" db:"full_name"`
	Email    string `json:"email" query:"email" db:"email"`
	Phone    string `json:"phone_number" query:"phonenumber" db:"phone_number"`
	Address  string `json:"address" query:"address" db:"address"`
	Gender   bool   `json:"gender" query:"gender" db:"gender"`
	JobTitle string `json:"job_title" query:"jobtitle" db:"job_title"`
	// DepartmentId int    `json:"department_id" query:"departmentid" db:"department_id"`
}
