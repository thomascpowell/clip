package main

import(
	"video-api/store"
	"video-api/utils"
	"video-api/workers"
	"video-api/server"
	"log"
	"os"
)

const WORKER_COUNT = 4

func main() {
	address := os.Getenv("REDIS_ADDR")
	if address == "" {
    log.Fatal("REDIS_ADDR env var is empty")
	}
	if err := store.InitRedis(address); err != nil {
		log.Fatalf("Error starting Redis: %v", err)
	}

	jobs := make(chan utils.Job, 100)
	for i := range WORKER_COUNT {
		go workers.Worker(i, jobs)
	}

	router := server.SetupRouter(jobs)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting Gin: %v", err)
	}
}
