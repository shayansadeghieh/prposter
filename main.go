package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shayansadeghieh/prposter/internal"
)

// Hello there
func main() {
	pr, err := internal.GhCommand()
	if err != nil {
		log.Fatalf("Error running gh command: %v. Does this branch have a PR open?", err)
	}

	namesToID, names, err := internal.ReadSlackUsers()

	// Check for len(names) to be 0 due to slack rate limits
	if err != nil {
		log.Fatalf("Error reading slack users: %v", err)
	}

	if len(names) == 0 {
		log.Fatalf("Zero slack users read. This is likely due to slack's rate limits. Wait a minute and then try again.")
	}

	reviewer, err := internal.ReviewerPrompt(fmt.Sprintf("\033[1mEnter a reviewer for PR #%s: \033[0m", fmt.Sprint(pr.Number)), names)
	if err != nil {
		log.Fatalf("Error reading reviewer: %v", err)
	}

	prDescription := internal.DescriptionPrompt(fmt.Sprintf("\033[1mProvide a description for PR #%s: \033[0m", fmt.Sprint(pr.Number)))

	// Create the message format for Slack
	messageFormat := `{"channel": "%s", "text": "(+%s/-%s) <%s|PR #%s>: %s <@%s>"}`

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("SLACK_CHANNEL_ID environment variable is not set")
	}
	message := fmt.Sprintf(messageFormat, channelID, fmt.Sprint(pr.Additions), fmt.Sprint(pr.Deletions), pr.URL, fmt.Sprint(pr.Number), prDescription, namesToID[reviewer])

	err = internal.SendSlackMessage(message)
	if err != nil {
		log.Fatalf("Error sending message to Slack: %v", err)
	}

	fmt.Println("\033[1;32mPR successfully posted to slack.\033[0m")

}
