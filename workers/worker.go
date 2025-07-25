package workers

import (
	"clip-api/utils"
	"clip-api/store"
	"path/filepath"
	"log"
)

func Worker(id int, jobs <-chan utils.Job) {
	for job := range jobs {
		log.Printf("Worker %d ... Started job", id)
		Process(id, job)
	}
}

func Process(id int, job utils.Job) {
	if abortIfCanceled(id, job) { return }
	store.UpdateJobStatus(job.ID, utils.StatusProcessing)
	err := utils.Dlp(
		job.ID+".mp4", // file name
		job.URL,
	)
	if err != nil {
		log.Printf("Worker %d ... Error in dlp(): %v", id, err)
		store.UpdateJobStatus(job.ID, utils.StatusError)
		job.ResponseChan <- utils.Result{
			OutputPath: "", 
			Err: err,
		}
		close(job.ResponseChan)
		return
	}
	if abortIfCanceled(id, job) { return }
	err = utils.FFmpeg(
		job.ID+".mp4", // input_file_name = job id + .mp4 
		job.ID, // output_base
		job.Format, // output_format
		job.VolumeScale,
		job.StartTime,
		job.EndTime,
	)
	if err != nil {
		log.Printf("Worker %d ... Error in ffmpeg(): %v", id, err)
		store.UpdateJobStatus(job.ID, utils.StatusError)
		job.ResponseChan <- utils.Result{
			OutputPath: "", 
			Err: err,
		}
		close(job.ResponseChan)
		return
	}
	if abortIfCanceled(id, job) { return }
	outPath := filepath.Join(utils.GetDir(), "out_" + job.ID + "." + job.Format)
	store.UpdateJobStatus(job.ID, utils.StatusDone)
	log.Printf("Worker %d ... Finished job", id)
	job.ResponseChan <- utils.Result{
		OutputPath: outPath, 
		Err: err,
	}
	close(job.ResponseChan)
}

func abortIfCanceled(id int, job utils.Job) bool {
	select {
	case <-job.Context.Done():
		// context expired or canceled
		job.ResponseChan <- utils.Result{
			OutputPath: "",
			Err: job.Context.Err(),
		}
		log.Printf("Worker %d ... Canceled job", id)
		close(job.ResponseChan)
		store.UpdateJobStatus(job.ID, utils.StatusError)
		return true
	default:
		return false
	}
}
