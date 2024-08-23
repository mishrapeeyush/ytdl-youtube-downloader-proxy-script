package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	// Define the YouTube video URL and desired audio format
	videoURL := "https://www.youtube.com/watch?v=p572p-irRaU&t=6s"
	audioFormat := "mp3"
	userAgent := "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36"

	// Generate a filename based on the current timestamp
	timestamp := time.Now().Format("20060102_150405")
	outputFilename := fmt.Sprintf("%s-%s.%s", "peeyush", timestamp, audioFormat)

	// Command to execute yt-dlp with the specified options
	cmd := exec.Command("yt-dlp", "--proxy", "http://hgfumqbe:t8a93hs9ef3r@45.127.248.127:5128", "-x", videoURL, "--user-agent", userAgent, "--audio-format", audioFormat, "-o", outputFilename)

	// Capture the output and error
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Command execution failed with error: %v", err)
	}

	// Print the output
	fmt.Println(string(output))
}
