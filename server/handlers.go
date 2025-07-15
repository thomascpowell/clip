package server

import(
	"video-api/utils"
	"video-api/store"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlePostVideo(jobs chan utils.Job) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var job utils.Job

		if err := ctx.ShouldBindJSON(&job); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		job.ID = uuid.New().String()
		job.Context = ctx.Request.Context()
		job.ResponseChan = make(chan utils.Result)

		if err := store.StoreJob(job.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to queue job"})
			return
		}


		// send to job queue
		utils.StartJob(jobs, job)


		ctx.JSON(http.StatusOK, gin.H{
			"message": "job accepted",
			"id":      job.ID,
		})
	}
}

func HandleGetVideo() error {
	return nil
}
