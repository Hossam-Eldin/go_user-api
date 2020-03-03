package users

import (
	"github.com/Hossam-Eldin/go_user-api/database/mysql/usersdb"
	"github.com/Hossam-Eldin/go_user-api/utils/date"
	"github.com/Hossam-Eldin/go_user-api/utils/errors"
	"github.com/Hossam-Eldin/go_user-api/utils/mysqlutils"
)

var (
	usersFakeDB = make(map[int64]*User)
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT * FROM users WHERE id=?;"
)

//Get : to find user by id
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysqlutils.ParseError(err)

	}
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
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err)

	}
	user.ID = userID
	return nil
}
