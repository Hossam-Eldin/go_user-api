package users

import (
	"net/http"

	"github.com/Hossam-Eldin/go_user-api/domain/users"
	"github.com/Hossam-Eldin/go_user-api/services"
	"github.com/Hossam-Eldin/go_user-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// CreateUser : in database
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

//SearchUser : find user
func SearchUser(c *gin.Context) {}

//GetUser : get user data
func GetUser(c *gin.Context) {}
