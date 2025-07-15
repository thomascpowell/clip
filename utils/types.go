package utils

import(
	"context"
)

type Job struct {
	Context      context.Context `json:"-"`
	ID           string          `json:"-"`
	URL          string          `json:"url" binding:"required"`
	Format       string          `json:"format"`
	VolumeScale  string          `json:"volumeScale"`
	StartTime    string          `json:"startTime"`
	EndTime      string          `json:"endTime"`
	ResponseChan chan Result     `json:"-"`
}

type Result struct {
	OutputPath  string
	Err         error
}

// WIP
// figuring out how to store job status
type JobStatus string
const (
	StatusQueued     JobStatus = "queued"
	StatusProcessing JobStatus = "processing"
	StatusDone       JobStatus = "done"
	StatusError      JobStatus = "error"
)
