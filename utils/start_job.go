package utils

import (
	"context"
	"time"
)

func StartJob(jobs chan Job, job Job) Result {
	job.ResponseChan = make(chan Result, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	job.Context = ctx
	jobs <- job
	select {
	case result := <-job.ResponseChan:
		// write status here?
		return result
	case <-ctx.Done():
		return Result{"", ctx.Err()}
	}
}
