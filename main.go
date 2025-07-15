package main

import(
	"video-api/store"
	"video-api/utils"
	"video-api/server"
	"log"
)

const WORKER_COUNT = 4

func main() {
	if err := store.InitRedis("localhost:6379"); err != nil {
		log.Fatalf("Error starting Redis: %v", err)
	}
	jobs := make(chan utils.Job, 100)

	for i := range WORKER_COUNT {
		go utils.Worker(i, jobs)
	}

	router := server.SetupRouter(jobs)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting Gin: %v", err)
	}
}

// test main

// import (
// 	"fmt"
// )
// func main() {
// 	utils.MakeDirectory()
// 	jobs := make(chan utils.Job, WORKER_COUNT) 
//

//
// 	job1 := utils.Job{
// 		ID:          uuid.New().String(),
// 		URL:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
// 		Format:      "mp3",
// 		VolumeScale: "1.5",
// 		StartTime:   "00:00:10",
// 		EndTime:     "00:00:20",
// 	}
//
// 	results := make(chan utils.Result, 2)
//
// 	go func() {
//     results <- utils.StartJob(jobs, job1)
// 	}()
//
// 	res1 := <-results
//
// 	fmt.Println("Result 1:", res1.OutputPath)
// }
