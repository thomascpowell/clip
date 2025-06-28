package main

import (
	"video-api/utils"
	"fmt"
)

const WORKER_COUNT = 2

func main() {
	jobs := make(chan utils.Job, WORKER_COUNT) 

	for i := range WORKER_COUNT {
		go utils.Worker(i, jobs)
	}

	job1 := utils.Job{
		ID:          "test1",
		URL:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		Format:      "mp3",
		VolumeScale: "1.5",
		StartTime:   "00:00:10",
		EndTime:     "00:00:20",
	}

	job2 := utils.Job{
		ID:          "test2",
		URL:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		Format:      "mp4",
		VolumeScale: "1",
		StartTime:   "00:00:00",
		EndTime:     "00:00:20",
	}

	result1 := utils.StartJob(jobs, job1)
	result2 := utils.StartJob(jobs, job2)

	if result1.Err != nil {
		fmt.Println("job1 failed:", result1.Err)
	} else {
		fmt.Println("success: ", result1.OutputPath)
	}

	if result2.Err != nil {
		fmt.Println("job2 failed:", result2.Err)
	} else {
		fmt.Println("success: ", result2.OutputPath)
	}

}
