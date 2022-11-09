package sql

import (
	"fmt"

	"github.com/oabraham1/mongosqlgen/internal/parser"
)

type Command string

const (
	SQLSelect Command = "SELECT"
	SQLInsert Command = "INSERT"
	SQLUpdate Command = "UPDATE"
	SQLDelete Command = "DELETE"
)

type Query struct {
	Command  Command
	Database string
	Table    string
	Columns  string
	Filter   string
	Values   interface{}
}

func ParseSQLCommand(command string) (Command, error) {
	switch command {
	case "SELECT":
		return SQLSelect, nil
	case "INSERT":
		return SQLInsert, nil
	case "UPDATE":
		return SQLUpdate, nil
	case "DELETE":
		return SQLDelete, nil
	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}
}

func GetCommandFromUserInput(input string) (Command, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return "", err
	}
	return ParseSQLCommand(tokens[0])
}

// GetColumns returns a single columns from a string of the form (col1)
func GetColumnsForInsert(input string) (string, error) {
	token, err := parser.SplitInputByDelimiters(input, []string{"(", ")"})
	if err != nil {
		return "", err
	}
	return token[0], nil
}

// GetValues returns a single value from a string of the form (val1)
func GetValuesForInsert(input string) (string, error) {
	token, err := parser.SplitInputByDelimiters(input, []string{"(", ")"})
	if err != nil {
		return "", err
	}
	return token[0], nil
}

// GetColumns returns a single columns from a string of the form col1=val1 or col1 = val1
func GetColumnsAndValuesForUpdate(input string) ([]string, error) {
	token, err := parser.SplitInputByDelimiters(input, []string{"=", " "})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CovertUserInputToSQLQuery(input string) (Query, error) {
	command, err := GetCommandFromUserInput(input)
	if err != nil {
		return Query{}, err
	}
	if command == SQLSelect {
		return HandleSelectUserInput(input)
	}
	if command == SQLInsert {
		return HandleInsertUserInput(input)
	}
	if command == SQLUpdate {
		return HandleUpdateUserInput(input)
	}
	if command == SQLDelete {
		return HandleDeleteUserInput(input)
	}
	return Query{}, fmt.Errorf("unknown command: %s", command)
}

// Can only select from one table and one column
func HandleSelectUserInput(input string) (Query, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return Query{}, err
	}
	if len(tokens) != 4 {
		return Query{}, fmt.Errorf("invalid input: %s", input)
	}
	return Query{
		Command: SQLSelect,
		Table:   tokens[3],
		Columns: tokens[1],
	}, nil
}

// Handle insert into one column
func HandleInsertUserInput(input string) (Query, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return Query{}, err
	}
	if len(tokens) != 6 {
		return Query{}, fmt.Errorf("invalid input: %s", tokens)
	}
	column, err := GetColumnsForInsert(tokens[3])
	if err != nil {
		return Query{}, err
	}
	value, err := GetValuesForInsert(tokens[5])
	if err != nil {
		return Query{}, err
	}
	return Query{
		Command: SQLInsert,
		Table:   tokens[2],
		Columns: column,
		Values:  value,
	}, nil

}

// Handle update for one column without spaces in column and value
func HandleUpdateUserInput(input string) (Query, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return Query{}, err
	}
	if len(tokens) != 6 {
		return Query{}, fmt.Errorf("invalid input: %s", input)
	}
	colandval, err := GetColumnsAndValuesForUpdate(tokens[3])
	if err != nil {
		return Query{}, err
	}
	return Query{
		Command: SQLUpdate,
		Table:   tokens[1],
		Columns: colandval[0],
		Values:  colandval[1],
		Filter:  tokens[5],
	}, nil
}

// Handle delete for one column without spaces in column and value
func HandleDeleteUserInput(input string) (Query, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return Query{}, err
	}
	if len(tokens) != 5 {
		return Query{}, fmt.Errorf("invalid input: %s", input)
	}
	return Query{
		Command: SQLDelete,
		Table:   tokens[2],
		Filter:  tokens[4],
	}, nil
}
