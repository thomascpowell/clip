package utils

import (
	"path/filepath"
	"time"
	"fmt"
)

func Worker(id int, jobs <-chan Job) {
	for job := range jobs {
		log := fmt.Sprintf("[%s] Worker %d / Job ID %s", time.Now().Format("15:04:05.000"), id, job.ID)
		fmt.Println(log)
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
		close(job.ResponseChan)
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
		close(job.ResponseChan)
		return
	}
	if isCanceled(job) { return }
	outPath := filepath.Join(GetDir(), "out_" + job.ID + "." + job.Format)
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

