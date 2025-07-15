package server

import (
	"github.com/gin-gonic/gin"
	"video-api/utils"
)

func SetupRouter(jobs chan<- utils.Job) *gin.Engine {
	router := gin.Default()
	router.POST("/videos", HandlePostVideo(jobs))
	router.GET("/videos/:id", HandleGetVideo)
	return router
}
