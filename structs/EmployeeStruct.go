package structs

type Employee struct {
	Id       int    `json:"employee_id"`
	Name     string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone_number"`
	Address  string `json:"address"`
	Gender   bool   `json:"gender"`
	JobTitle string `json:"job_title"`
	// DepartmentId int    `json:"department_id" query:"departmentid" db:"department_id"`
}
