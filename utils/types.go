package utils

import(
	"context"
)

type Job struct {       	
	Context      context.Context
	ID           string
	URL          string
	Format       string
	VolumeScale  string
	StartTime    string
	EndTime      string
	ResponseChan chan Result
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
