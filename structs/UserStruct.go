package structs

type User struct {
	Id    string `json:"user_id" query:"userid" db:"id"`
	Name  string `json:"user_name" query:"username" db:"name"`
	Email string `json:"personal_email" query:"useremail" db:"email"`
	Phone string `json:"phone_number" query:"phonenumber" db:"phone"`
}
