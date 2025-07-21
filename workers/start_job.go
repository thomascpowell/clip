package workers

import (
	"clip-api/utils"
	"context"
	"time"
)

func StartJob(jobs chan utils.Job, job utils.Job) utils.Result {
	job.ResponseChan = make(chan utils.Result, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	job.Context = ctx
	jobs <- job
	select {
	case result := <-job.ResponseChan:
		// write status here?
		return result
	case <-ctx.Done():
		return utils.Result{"", ctx.Err()}
	}
}
