package utils

import (
	"path/filepath"
	"os"
)

func Worker(id int, jobs <-chan Job) {
	for job := range jobs {
		println("Worker: ", id, "... Started Job: ", job.ID)
		Process(id, job)
	}
}

func Process(id int, job Job) {
	if isCanceled(job) { return }
	err := dlp(
		job.ID+".mp4", // file name
		job.URL,
	)
	if err != nil {
		println("Worker: ", id, "... Error in dlp(): ", err.Error())
		job.ResponseChan <- Result{"", err}
		return
	}
	if isCanceled(job) { return }
	err = ffmpeg(
		job.ID+".mp4", // input_file_name = job id + .mp4 
		job.ID, // output_base
		job.Format, // output_format
		job.VolumeScale,
		job.StartTime,
		job.EndTime,
	)
	if err != nil {
		println("Worker: ", id, "... Error in ffmpeg(): ", err.Error())
		job.ResponseChan <- Result{"", err}
		return
	}
	outPath := filepath.Join(os.TempDir(), "out_" + job.ID + "." + job.Format)
	job.ResponseChan <- Result{outPath, nil}
	close(job.ResponseChan)
}

func isCanceled(job Job) bool {
	select {
	case <-job.Context.Done():
		// context expired or canceled
		job.ResponseChan <- Result{"", job.Context.Err()}
		close(job.ResponseChan)
		return true
	default:
		return false
	}
}

