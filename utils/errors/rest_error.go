package errors

import (
	"errors"
	"net/http"
)

//RestErr : struct error
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

//NewError : error new
func NewError(msg string) error {
	return errors.New(msg)
}

//NewBadRequestError : to handle the same type of error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError : to handle the same type of error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_Found",
	}
}

//NewInternalServerError : to hanlde db erros
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "insternal_server_error",
	}
}
