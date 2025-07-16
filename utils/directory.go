package utils

import (
	"os"
	"fmt"
)

func MakeDirectory() {
	err := os.MkdirAll(GetDir(), 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create temp directory", err.Error())
		os.Exit(1)
	}
}

func GetDir() string {
    return "/app/videos"
}
