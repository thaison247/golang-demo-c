package structs

type User struct {
	Id    int    `json:"-" query:"userid" db:"employee_id"`
	Name  string `json:"user_name" query:"username" db:"full_name"`
	Email string `json:"personal_email" query:"useremail" db:"email"`
	Phone string `json:"phone_number" query:"phonenumber" db:"phone_number"`
}
