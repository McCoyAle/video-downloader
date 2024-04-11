package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func main() {
	// Your YouTube API key
	apiKey := "YOUR_API_KEY"

	// Create a new YouTube client
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		fmt.Println("Error creating YouTube client:", err)
		return
	}

	// ID of the video you want to download
	videoID := "VIDEO_ID"

	// Get video details
	videoCall := service.Videos.List("snippet").Id(videoID)
	response, err := videoCall.Do()
	if err != nil {
		fmt.Println("Error fetching video details:", err)
		return
	}

	// Extract video title
	videoTitle := response.Items[0].Snippet.Title
	fmt.Println("Downloading:", videoTitle)

	// Download video
	resp, err := http.Get(fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID))
	if err != nil {
		fmt.Println("Error downloading video:", err)
		return
	}
	defer resp.Body.Close()

	// Create output file
	file, err := os.Create(videoTitle + ".mp4")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	// Copy video content to output file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving video:", err)
		return
	}

	fmt.Println("Video downloaded successfully!")
}
