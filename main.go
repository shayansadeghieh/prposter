package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// Set environment variables
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		log.Fatal("SLACK_API_TOKEN environment variable is not set")
	}
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("SLACK_CHANNEL_ID environment variable is not set")
	}

	// TODO: Read in the user's message
	messageFormat := `{"channel": "%s", "text": "Hello, World!"}`
	message := fmt.Sprintf(messageFormat, channelID)
	requestData := []byte(message)

	// Initialize a request to the slack post message API
	slackUrl := "https://slack.com/api/chat.postMessage"
	request, err := http.NewRequest("POST", slackUrl, bytes.NewBuffer(requestData))
	if err != nil {
		log.Fatalf("Error initializing NewRequest: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token)) // Set the token in the Authorization header

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error sending request to %s: %v", slackUrl, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	println(string(body))
}
