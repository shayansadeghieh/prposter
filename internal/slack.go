package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func extractNames(members AllMembers) []string {
	names := make([]string, len(members.Members))
	for i, member := range members.Members {
		names[i] = member.Profile.RealNameNormalized
	}
	return names
}

type AllMembers struct {
	Members []Member `json:"members"`
}

type Member struct {
	Profile Profile `json:"profile"`
}

type Profile struct {
	// Add other fields if needed
	RealNameNormalized string `json:"real_name_normalized"`
}

func ReadSlackUsers() ([]string, error) {
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("SLACK_API_TOKEN environment variable is not set")
	}
	url := "https://slack.com/api/users.list"

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var slackMembers AllMembers
	err = json.Unmarshal([]byte(body), &slackMembers)
	if err != nil {
		log.Fatalf("Error unmarshaling output from slackMembers: %v", err)
	}

	// Extract names and place them into a slice of names
	names := extractNames(slackMembers)

	return names, nil

}

func SendSlackMessage(message string) error {
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
