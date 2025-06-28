package utils

import(
	"os"
	"fmt"
	"os/exec"
	"path/filepath"
)

/**
* Wrappers for really good CLI tools.
*/

/**
* naming convention
* file_name: e.g. video.mp4
* format:    e.g. mp4
* base:      e.g. video
*/

func dlp(output_file_name string, url string) error {
	directory := GetDir()
	path := filepath.Join(directory, output_file_name)
	cmd := exec.Command("yt-dlp", 
		"-o", path, 
		"-f", "mp4",
		url)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// var stderr bytes.Buffer
	// cmd.Stderr = &stderr
	return cmd.Run()
}

func ffmpeg(input_file_name string, output_base string, output_format string, volume_scale string, start_time string, end_time string) error {
	directory := GetDir()
	input_path := filepath.Join(directory, input_file_name)
	output_path := filepath.Join(directory, "out_"+output_base+ "." + output_format)
	
	args := []string{}

	// options
	if start_time != "" {
		args = append(args, "-ss", start_time)
	}
	args = append(args, "-i", input_path)
	if end_time != "" {
		args = append(args, "-to", end_time)
	}
	if volume_scale != "" {
		args = append(args, "-af", fmt.Sprintf("volume=%s", volume_scale))
	}

	// format encoders
	switch output_format {
	case "mp3":
		args = append(args, "-vn", "-c:a", "libmp3lame")
	case "wav":
		args = append(args, "-vn", "-c:a", "pcm_s16le")
	case "mp4":
		args = append(args, "-c:v", "copy", "-c:a", "aac")
	default:
		return fmt.Errorf("unsupported format: %s", output_format)
	}

	args = append(args, output_path)
	cmd := exec.Command("ffmpeg", args...)
	// var stderr bytes.Buffer
	// cmd.Stderr = &stderr
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	defer os.Remove(input_path)
	return cmd.Run()
}
