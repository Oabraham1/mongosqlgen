package parser

import "fmt"

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

func contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
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
