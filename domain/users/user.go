package users

//User : user in database
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firsT_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}
