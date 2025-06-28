package main

import (
	"video-api/utils"
	"fmt"
)

const WORKER_COUNT = 2

func main() {
	utils.MakeDirectory()
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

	results := make(chan utils.Result, 2)

	go func() {
    results <- utils.StartJob(jobs, job1)
	}()

	res1 := <-results
	
	fmt.Println("Result 1:", res1.OutputPath)
}
