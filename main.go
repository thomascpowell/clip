package main

import (
	"video-api/utils"
)

const WORKER_COUNT = 2

func main() {
	jobs := make(chan utils.Job, WORKER_COUNT) 
}
