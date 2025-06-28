package utils

import(
	"os"
	"path/filepath"
)

type Job struct {
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

func Worker(id int, jobs <-chan Job) {
	for job := range jobs {
		println("Worker: ", id, "... Started Job: ", job.ID)
		process(id, job)
	}
}

func process(id int, job Job) {
	err := dlp(
		job.ID+".mp4", // file name
		job.URL,
	)
	if err != nil {
		println("Worker: ", id, "... Error in dlp(): ", err.Error())
		job.ResponseChan <- Result{"", err}
		return
	}
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
