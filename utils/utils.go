package utils

import (
	"fmt"
	"os"
	"strings"
)

func NotAllSpaces(s string) bool { return strings.TrimSpace(s) != "" }

func AskUser(prompt string, isValid func(string) bool) string {
	fmt.Print(prompt)
	var userInput string
	if _, err := fmt.Scanln(&userInput); err != nil || !isValid(userInput) {
		fmt.Println("Quitting.")
		os.Exit(1)
	}
	return userInput
}
