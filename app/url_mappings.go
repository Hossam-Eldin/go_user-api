package app

import (
	users "github.com/Hossam-Eldin/go_user-api/controllers"
	"github.com/Hossam-Eldin/go_user-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
}
