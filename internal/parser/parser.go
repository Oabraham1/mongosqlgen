package parser

import (
	"fmt"
	"strings"
)

// ParseUserInput parses user input into a slice of tokens
func ParseUserInput(input string) ([]string, error) {
	var tokens []string
	var token string
	for _, char := range input {
		if char == ' ' {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
		} else {
			token += string(char)
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid input: %s", input)
	}
	return tokens, nil
}

// FindAfter finds the first instance of a string and returns the string after it
func FindAfter(input string, s string) (string, error) {
	size := len(s) - 1
	index := strings.Index(input, s) + size
	if index == -1 {
		return "", fmt.Errorf("could not find: %s", s)
	}
	return input[index+1:], nil
}

// FindBetween finds the first instance of a string and returns the string between it
func FindBetween(input string, start string, end string) (string, error) {
	size := len(start) - 1
	startIndex := strings.Index(input, start) + size
	if startIndex == -1 {
		return "", fmt.Errorf("could not find start: %s", start)
	}
	endIndex := strings.Index(input, end)
	if endIndex == -1 {
		return "", fmt.Errorf("could not find end: %s", end)
	}
	return input[startIndex+1 : endIndex], nil
}

// ContainsCommand checks if a string contains a command
func ContainsCommand(input string, s string) bool {
	tokens := strings.Split(input, " ")
	for _, token := range tokens {
		if token == s {
			return true
		}
	}
	return false
}

// contains checks if a string slice contains a string
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// SplitInputByDelimiters splits a string by a slice of delimiters
func SplitInputByDelimiters(input string, delimiters []string) ([]string, error) {
	var tokens []string
	var token string
	for _, char := range input {
		if contains(delimiters, string(char)) {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
		} else {
			token += string(char)
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	if len(tokens) < 1 {
		return nil, fmt.Errorf("invalid input: %s", input)
	}
	return tokens, nil
}

// CountOccurences counts the number of times a string appears in another string
func CountOccurences(input string, s string) int {
	count := 0
	for _, char := range input {
		if string(char) == s {
			count++
		}
	}
	return count
}
