package main

import (
	"fmt"
	"log"
	"os/exec"
)

// func prPrompt(label string) string {
// 	var s string
// 	r := bufio.NewReader(os.Stdin)
// 	// for {
// 		cmd := exec.Command("gh", "pr", "list")
// 		// fmt.Fprint(os.Stderr, label+": ")

// 		if err := cmd.Run(); err != nil {
// 			log.Fatal(err)
// 		}

// 		s, _ = r.ReadString('\n')
// 		if s != "" {
// 			break
// 		}
// 	}
// 	return strings.TrimSpace(s)

// }

func main() {
	cmd := exec.Command("gh", "pr", "list")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running gh pr list command: %v", err)
	} else {
		fmt.Println("Output:", string(output))
	}

	// pr := prPrompt("Choose the PR to post")
	// fmt.Println("the pr is", pr)

	// Set environment variables
	// token := os.Getenv("SLACK_API_TOKEN")
	// if token == "" {
	// 	log.Fatal("SLACK_API_TOKEN environment variable is not set")
	// }
	// channelID := os.Getenv("SLACK_CHANNEL_ID")
	// if channelID == "" {
	// 	log.Fatal("SLACK_CHANNEL_ID environment variable is not set")
	// }

	// // TODO: Read in the user's message
	// messageFormat := `{"channel": "%s", "text": "Hello, World!"}`
	// message := fmt.Sprintf(messageFormat, channelID)
	// requestData := []byte(message)

	// // Initialize a request to the slack post message API
	// slackUrl := "https://slack.com/api/chat.postMessage"
	// request, err := http.NewRequest("POST", slackUrl, bytes.NewBuffer(requestData))
	// if err != nil {
	// 	log.Fatalf("Error initializing NewRequest: %v", err)
	// }
	// request.Header.Set("Content-Type", "application/json")
	// request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token)) // Set the token in the Authorization header

	// client := http.Client{}
	// response, err := client.Do(request)
	// if err != nil {
	// 	log.Fatalf("Error sending request to %s: %v", slackUrl, err)
	// }
	// defer response.Body.Close()

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatalf("Error reading response body: %v", err)
	// }

	// println(string(body))
}
