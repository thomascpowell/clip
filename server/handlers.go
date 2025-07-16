package server

import(
	"path/filepath"
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
		if err := store.StoreJob(job.ID, job.Format); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to queue job"})
			return
		}
		utils.StartJob(jobs, job) // sends to job queue
		ctx.JSON(http.StatusOK, gin.H{
			"message": "job accepted",
			"id":      job.ID,
		})
	}
}

func HandleGetVideo(ctx *gin.Context) {
	id := ctx.Param("id")
	status, format, err := store.GetStatusAndFormat(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
		return
	}
	switch status {
	case utils.StatusDone:
		outputPath := filepath.Join(utils.GetDir(), "out_" + id + "." + format)
		ctx.File(outputPath)
	case utils.StatusError:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "job failed or was canceled"})
	case utils.StatusQueued, utils.StatusProcessing:
		ctx.JSON(http.StatusAccepted, gin.H{"message": "job is still processing"})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unknown job status"})
	}
}
