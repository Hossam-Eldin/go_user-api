package users

import (
	"strings"

	"github.com/Hossam-Eldin/go_user-api/utils/errors"
)

//User : user in database
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firsT_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate : for user data
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
