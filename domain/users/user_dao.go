package users

import (
	"fmt"

	"github.com/Hossam-Eldin/go_user-api/database/mysql/usersdb"
	"github.com/Hossam-Eldin/go_user-api/logger"
	"github.com/Hossam-Eldin/go_user-api/utils/errors"
)

var (
	usersFakeDB = make(map[int64]*User)
)

const (
	errorNoRows           = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email,  status, password,date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT * FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?,  email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT * FROM users WHERE status=?;"
)

//Get : to find user by id
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error When trying to prepare get user statment", err)
		return errors.NewInternalServerError("error when trying to get user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.Password, &user.DateCreated); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//Save : to insert user data
func (user *User) Save() *errors.RestErr {

	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepre save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated)
	if saveErr != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("database error")

	}
	user.ID = userID
	return nil
}

//Update : update user data
func (user *User) Update() *errors.RestErr {

	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error when trying to update user", err)

		return errors.NewInternalServerError("database error")
	}

	return nil
}

// Delete : delete method
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)

		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")

	}
	return nil
}

//FindByStatus : find user by status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find users by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find users by status", err)

		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			logger.Error("error when scan user row into user struct", err)

			return nil, errors.NewInternalServerError("database error")
		}

		results = append(results, user)
	}

	if len(results) == 0 {

		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
