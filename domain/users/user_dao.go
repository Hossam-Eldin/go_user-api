package users

import (
	"fmt"
	"strings"

	"github.com/Hossam-Eldin/go_user-api/database/mysql/usersdb"
	"github.com/Hossam-Eldin/go_user-api/utils/date"
	"github.com/Hossam-Eldin/go_user-api/utils/errors"
)

var (
	usersFakeDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

//Get : to find user by id
func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersFakeDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.ID))
	}
	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

//Save : to insert user data
func (user *User) Save() *errors.RestErr {

	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email") {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s Already Exist ", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save the user %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save the user %s", err.Error()))

	}
	user.ID = userID
	return nil
}
