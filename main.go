package main

import (
	"fmt"
	"log"
	"os"
	"prposter/internal"
)

func main() {
	names, err := internal.ReadSlackUsers()
	if err != nil {
		log.Fatalf("Error reading slack users: %v", err)
	}
	fmt.Println("names are", names)

	pr, err := internal.GhCommand()
	if err != nil {
		log.Fatalf("Error running gh command: %v", err)
	}
	fmt.Println("pr is", pr)

	prDescription := internal.StringPrompt(fmt.Sprintf("Provide a description for PR #%s:", fmt.Sprint(pr.Number)))

	// Create the message format for Slack
	messageFormat := `{"channel": "%s", "text": "(+%s/-%s) <%s|PR #%s>: %s"}`

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("SLACK_CHANNEL_ID environment variable is not set")
	}
	message := fmt.Sprintf(messageFormat, channelID, fmt.Sprint(pr.Additions), fmt.Sprint(pr.Deletions), pr.URL, fmt.Sprint(pr.Number), prDescription)

	err = internal.SendSlackMessage(message)
	if err != nil {
		log.Fatalf("Error sending message to Slack: %v", err)
	}

}
