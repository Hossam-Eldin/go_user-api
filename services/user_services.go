package services

import (
	"github.com/Hossam-Eldin/go_user-api/domain/users"
	"github.com/Hossam-Eldin/go_user-api/utils/errors"
)

//CreateUser : service function
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
