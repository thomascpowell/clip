package workers

import (
	"video-api/utils"
	"video-api/store"
	"path/filepath"
	"time"
	"fmt"
)

func Worker(id int, jobs <-chan utils.Job) {
	for job := range jobs {
		log := fmt.Sprintf("[%s] Worker %d / utils.Job ID %s", time.Now().Format("15:04:05.000"), id, job.ID)
		fmt.Println(log)
		Process(id, job)
	}
}

func Process(id int, job utils.Job) {
	if abortIfCanceled(job) { return }
	store.UpdateJobStatus(job.ID, utils.StatusProcessing)
	err := utils.Dlp(
		job.ID+".mp4", // file name
		job.URL,
	)
	if err != nil {
		println("Worker: ", id, "... Error in dlp(): ", err.Error())
		job.ResponseChan <- utils.Result{
			OutputPath: "", 
			Err: err,
		}
		close(job.ResponseChan)
		return
	}
	if abortIfCanceled(job) { return }
	err = utils.FFmpeg(
		job.ID+".mp4", // input_file_name = job id + .mp4 
		job.ID, // output_base
		job.Format, // output_format
		job.VolumeScale,
		job.StartTime,
		job.EndTime,
	)
	if err != nil {
		println("Worker: ", id, "... Error in ffmpeg(): ", err.Error())
		job.ResponseChan <- utils.Result{
			OutputPath: "", 
			Err: err,
		}
		close(job.ResponseChan)
		return
	}
	if abortIfCanceled(job) { return }
	outPath := filepath.Join(utils.GetDir(), "out_" + job.ID + "." + job.Format)
	store.UpdateJobStatus(job.ID, utils.StatusDone)
	job.ResponseChan <- utils.Result{
		OutputPath: outPath, 
		Err: err,
	}
	close(job.ResponseChan)
}

func abortIfCanceled(job utils.Job) bool {
	select {
	case <-job.Context.Done():
		// context expired or canceled
		job.ResponseChan <- utils.Result{
			OutputPath: "",
			Err: job.Context.Err(),
		}
		close(job.ResponseChan)
		store.UpdateJobStatus(job.ID, utils.StatusError)
		return true
	default:
		return false
	}
}

