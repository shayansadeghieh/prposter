package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type PullRequest struct {
	URL string `json:"url"`
}

func extractPrNumber(url string) string {
	urlPortions := strings.Split(url, "/")
	return urlPortions[len(urlPortions)-1]
}

func sendSlackMessage(message string) error {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		return fmt.Errorf("SLACK_API_TOKEN environment variable is not set")
	}

	// Slack API URL
	url := "https://slack.com/api/chat.postMessage"

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack API request failed with status: %s", resp.Status)
	}

	return nil
}

func main() {
	cmd := exec.Command("gh", "pr", "view", "--json", "url")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running gh pr view command: %v", err)
	}

	var pr PullRequest
	err = json.Unmarshal([]byte(output), &pr)

	if err != nil {
		log.Fatalf("Error unmarshaling output from gh pr view command: %v", err)
	}

	prNumber := extractPrNumber(pr.URL)

	// Create the message format for Slack
	messageFormat := `{"channel": "%s", "text": "<%s|PR #%s>"}`

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("SLACK_CHANNEL_ID environment variable is not set")
	}
	message := fmt.Sprintf(messageFormat, channelID, pr.URL, prNumber)

	err = sendSlackMessage(message)
	if err != nil {
		log.Fatalf("Error sending message to Slack: %v", err)
	}
}
