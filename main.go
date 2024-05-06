package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"prposter/internal"
	"strings"
)

func filterNames(names []string, filter string) []string {
	var filteredNames []string
	for _, name := range names {
		if strings.Contains(strings.ToLower(name), strings.ToLower(filter)) {
			filteredNames = append(filteredNames, name)
		}
	}

	return filteredNames
}

func StringPromptReview(prompt string, names []string) string {
	reader := bufio.NewReader(os.Stdin)

	var reviewer string
	var counter int
	for {
		if counter > 0 {
			fmt.Print("\nEnter a reviewer (we fuzzy match): ")
		} else {
			fmt.Print("Enter a reviewer (we fuzzy match): ")
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		// Trim newline character
		input = strings.TrimSpace(input)

		// If input is empty, exit
		if input == "" {
			break
		}

		// Filter and print names based on input
		filteredNames := filterNames(names, input)

		// If we receive more than one name, say something witty prompt the user to choose one
		// using the number next to the name
		// Switch on the length of filtered names

		if len(filteredNames) == 0 {
			fmt.Println("0 results. Are you sure this person works here? Try again.")
		} else if len(filteredNames) == 1 {
			fmt.Println(filteredNames[0])
			reviewer = filteredNames[0]
			break
		} else {
			fmt.Printf("Too lazy. I found %d results. Try again.", len(filteredNames))
		}

		counter += 1

	}

	return reviewer
}

func main() {
	pr, err := internal.GhCommand()
	if err != nil {
		log.Fatalf("Error running gh command: %v. Does this branch have a PR open?", err)
	}

	names, err := internal.ReadSlackUsers()
	if err != nil {
		log.Fatalf("Error reading slack users: %v", err)
	}

	reviewer := StringPromptReview(fmt.Sprintf("Enter a reviewer for PR #%s: ", fmt.Sprint(pr.Number)), names)

	prDescription := internal.StringPrompt(fmt.Sprintf("Provide a description for PR #%s: ", fmt.Sprint(pr.Number)))

	// Create the message format for Slack
	messageFormat := `{"channel": "%s", "text": "(+%s/-%s) <%s|PR #%s>: %s @%s"}`

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("SLACK_CHANNEL_ID environment variable is not set")
	}
	message := fmt.Sprintf(messageFormat, channelID, fmt.Sprint(pr.Additions), fmt.Sprint(pr.Deletions), pr.URL, fmt.Sprint(pr.Number), prDescription, reviewer)

	err = internal.SendSlackMessage(message)
	if err != nil {
		log.Fatalf("Error sending message to Slack: %v", err)
	}
}
