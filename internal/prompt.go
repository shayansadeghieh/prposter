package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
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

func StringPrompt(prompt string) string {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[31m»\033[0m ",
		HistoryFile:       "/tmp/readline.tmp",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		line, _ := l.Readline()
		fmt.Println(line)
	}

	return ""
}

// func StringPrompt(prompt string) string {
// 	reader := bufio.NewReader(os.Stdin)
// 	var input string

// 	fmt.Print(prompt)

// 	for {
// 		char, _, err := reader.ReadRune()
// 		if err != nil {
// 			fmt.Println("Error reading input:", err)
// 			os.Exit(1)
// 		}

// 		if char == '\n' {
// 			break
// 		} else if char == '\b' { // Handle backspace
// 			if len(input) > 0 {
// 				input = input[:len(input)-1]
// 			}
// 		} else {
// 			input += string(char)
// 		}
// 	}

// 	return strings.TrimSpace(input)
// }