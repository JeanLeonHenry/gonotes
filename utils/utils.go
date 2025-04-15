package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	fzf "github.com/ktr0731/go-fuzzyfinder"
)

func NotAllSpaces(s string) bool { return strings.TrimSpace(s) != "" }

// AskUser prompts the user for input and calls os.Exit if isn't valid
func AskUser(prompt string, isValid func(string) bool) string {
	fmt.Print(prompt)
	var userInput string
	if _, err := fmt.Scanln(&userInput); err != nil || !isValid(userInput) {
		log.Fatalln("Quitting.")
	}
	return userInput
}

func AskUserAndSuggest(prompt string, isValid func(string) bool, suggestions []string) (string, error) {
	fmt.Print(prompt)
	choiceIndex, err := fzf.Find(suggestions, func(i int) string { return suggestions[i] })
	choice := fmt.Sprint(suggestions[choiceIndex])
	if err != nil || !isValid(choice) {
		return "", err
	}
	return choice, nil
}

func AllParsableIntoFloats(input []string) (string, int, error) {
	for k, s := range input {
		if _, err := strconv.ParseFloat(strings.Replace(s, ",", ".", 1), 64); err != nil {
			return s, k, err
		}
	}
	return "", 0, nil
}

func AllMatch(input []string, pattern string) (string, int, error) {
	for k, s := range input {
		match, err := regexp.MatchString(pattern, s)
		if err != nil {
			return s, k, err
		}
		if !match {
			return s, k, fmt.Errorf("No match")
		}
	}
	return "", 0, nil
}
