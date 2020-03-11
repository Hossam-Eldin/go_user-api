package app

import (
	"github.com/Hossam-Eldin/go_user-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Log.Info("about to start the app")
	router.Run(":5000")
}
