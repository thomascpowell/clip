package server

import (
	"github.com/gin-gonic/gin"
	"video-api/utils"
	"github.com/gin-contrib/cors"
	"time"
	"os"
)

func SetupRouter(jobs chan utils.Job) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	router.POST("/videos", HandlePostVideo(jobs))
	router.GET("/videos/:id", HandleGetVideo)
	return router
}
