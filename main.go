package gocreds

import (
	"errors"
	"fmt"
	"syscall"

	"golang.org/x/term"
)

// Prompt prompts user for non-sensitive data
func Prompt(prefix string) (string, error) {
	if len(prefix) > 0 {
		fmt.Printf("%s: ", prefix)
	} else {
		fmt.Print("Enter: ")
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", err
	}
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	return input, nil
}

// Prompt prompts user for sensitive data
func PromptSecret(prefix string) (string, error) {
	if len(prefix) > 0 {
		fmt.Printf("%s: ", prefix)
	} else {
		fmt.Print("Enter secret: ")
	}
	secret, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return "", err
	}
	if len(string(secret)) == 0 {
		return "", errors.New("empty secret")
	}
	fmt.Println()

	return string(secret), nil
}
