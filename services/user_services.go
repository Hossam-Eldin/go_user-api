package services

import (
	"github.com/Hossam-Eldin/go_user-api/domain/users"
	"github.com/Hossam-Eldin/go_user-api/utils/crypto"
	"github.com/Hossam-Eldin/go_user-api/utils/date"
	"github.com/Hossam-Eldin/go_user-api/utils/errors"
)

//GetUser : get user by id
func GetUser(userID int64) (*users.User, *errors.RestErr) {

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CreateUser : service function
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDbFormat()
	user.Password = crypto.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser : update user services
func UpdateUser(isPartail bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	if isPartail {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser : delete user services
func DeleteUser(userID int64) *errors.RestErr {
	user, err := GetUser(userID)
	if err != nil {
		return err
	}

	user = &users.User{ID: userID}
	return user.Delete()
}

//Search : to service status
func Search(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
