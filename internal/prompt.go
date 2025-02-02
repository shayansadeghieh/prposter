package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
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

func handleMultipleReviewers(reviewers []string) (string, error) {
	label := "\033[1mReviewer chosen: \033[0m"
	prompt := promptui.Select{
		Label: "Choose your reviewer",
		Items: reviewers,
		Templates: &promptui.SelectTemplates{
			Selected: fmt.Sprintf(`%s: {{ . | faint }}`, label),
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func ReviewerPrompt(prompt string, names []string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	var reviewer string

	for {
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		// Trim newline character
		input = strings.TrimSpace(input)

		// Filter and print names based on input
		filteredNames := filterNames(names, input)

		if len(filteredNames) == 0 {
			fmt.Println("0 results. Are you sure this person works here? Try again.")
		} else if len(filteredNames) == 1 {
			fmt.Printf("\033[1mReviewer chosen: \033[0m %s \n", filteredNames[0])
			reviewer = filteredNames[0]
			break
		} else {
			// If we receive more than one name, prompt the user to choose one
			reviewer, err := handleMultipleReviewers(filteredNames)
			if err != nil {
				return "", err
			}
			return reviewer, nil
		}

	}

	return reviewer, nil
}

func DescriptionPrompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	var input string

	fmt.Print(prompt)

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		if char == '\n' {
			break
		} else if char == '\b' { // Handle backspace
			if len(input) > 0 {
				input = input[:len(input)-1]
			}
		} else {
			input += string(char)
		}
	}

	return strings.TrimSpace(input)
}
