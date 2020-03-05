package users

import (
	"strings"

	"github.com/Hossam-Eldin/go_user-api/utils/errors"
)

const (
	//StatusActive :user status
	StatusActive = "active"
)

//User : user in database
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firsT_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Validate : for user data
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid Password")

	}

	return nil
}
