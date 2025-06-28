package utils

import (
	"os"
	"fmt"
	"path/filepath"
)

func MakeDirectory() {
	err := os.MkdirAll(GetDir(), 0755)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create temp directory", err.Error())
		os.Exit(1)
	}
}

func GetDir() string {
	tmp := os.TempDir()
	return filepath.Join(tmp, "video_api_tmp")
}
