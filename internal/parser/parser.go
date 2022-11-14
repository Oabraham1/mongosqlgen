package parser

import (
	"fmt"
	"strings"
)

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

func FindAfter(input string, s string) (string, error) {
	size := len(s) - 1
	index := strings.Index(input, s) + size
	if index == -1 {
		return "", fmt.Errorf("could not find: %s", s)
	}
	return input[index+1:], nil
}

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

func ContainsCommand(input string, s string) bool {
	tokens := strings.Split(input, " ")
	for _, token := range tokens {
		if token == s {
			return true
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

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

func CountOccurences(input string, s string) int {
	count := 0
	for _, char := range input {
		if string(char) == s {
			count++
		}
	}
	return count
}
