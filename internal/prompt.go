package internal

import (
	"fmt"
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

// Prompt for the reviewer.
func ReviewerPrompt(label string, names []string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	filteredNames := filterNames(names, result)

	if len(filteredNames) == 0 {
		fmt.Println("0 results. Are you sure this person works here?")
	} else if len(filteredNames) == 1 {
		return filteredNames[0], nil
	} else {
		reviewer, err := handleMultipleReviewers(filteredNames)
		if err != nil {
			return "", err
		}
		return reviewer, nil
	}

	return "", nil
}

func handleMultipleReviewers(reviewers []string) (string, error) {
	label := "Choose your reviewer"
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

// Prompt the user for a PR description
func DescriptionPrompt(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
