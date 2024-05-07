package internal

import (
	"encoding/json"
	"log"
	"os/exec"
)

type PullRequest struct {
	URL       string `json:"url"`
	Number    int    `json:"number"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
}

func GhCommand() (PullRequest, error) {
	cmd := exec.Command("gh", "pr", "view", "--json", "url,number,additions,deletions")
	output, err := cmd.Output()
	if err != nil {
		return PullRequest{}, err
	}

	var pr PullRequest
	err = json.Unmarshal([]byte(output), &pr)
	if err != nil {
		log.Fatalf("Error unmarshaling output from gh pr view command: %v", err)
	}

	return pr, nil
}
