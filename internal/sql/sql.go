package sql

import "fmt"
import "github.com/oabraham1/mongosqlgen/internal/parser"

type SQLCommand string

const (
	SQLSelect SQLCommand = "SELECT"
	SQLInsert SQLCommand = "INSERT"
	SQLUpdate SQLCommand = "UPDATE"
	SQLDelete SQLCommand = "DELETE"
)

type SQLQuery struct {
	Command  SQLCommand
	Database string
	Table    string
	Columns  string
	Filter   string
	Values   interface{}
}

func ParseSQLCommand(command string) (SQLCommand, error) {
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

func GetCommandFromUserInput(input string) (SQLCommand, error) {
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


func CovertUserInputToSQLQuery(input string) (SQLQuery, error) {
	command, err := GetCommandFromUserInput(input)
	if err != nil {
		return SQLQuery{}, err
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
	return SQLQuery{}, fmt.Errorf("unknown command: %s", command)
}

// Can only select from one table and one column
func HandleSelectUserInput(input string) (SQLQuery, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return SQLQuery{}, err
	}
	if len(tokens) != 4 {
		return SQLQuery{}, fmt.Errorf("invalid input: %s", input)
	}
	return SQLQuery{
		Command:  SQLSelect,
		Table:    tokens[3],
		Columns:  tokens[1],
	}, nil
}

// Handle insert into one column
func HandleInsertUserInput(input string) (SQLQuery, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return SQLQuery{}, err
	}
	if len(tokens) != 6 {
		return SQLQuery{}, fmt.Errorf("invalid input: %s", tokens)
	}
	column, err := GetColumnsForInsert(tokens[3])
	if err != nil {
		return SQLQuery{}, err
	}
	value, err := GetValuesForInsert(tokens[5])
	if err != nil {
		return SQLQuery{}, err
	}
	return SQLQuery{
		Command:  SQLInsert,
		Table:    tokens[2],
		Columns:  column,
		Values:   value,
	}, nil

}

// Handle update for one column without spaces in column and value
func HandleUpdateUserInput(input string) (SQLQuery, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return SQLQuery{}, err
	}
	if len(tokens) != 6 {
		return SQLQuery{}, fmt.Errorf("invalid input: %s", input)
	}
	colandval, err := GetColumnsAndValuesForUpdate(tokens[3])
	if err != nil {
		return SQLQuery{}, err
	}
	return SQLQuery{
		Command:  SQLUpdate,
		Table:    tokens[1],
		Columns:  colandval[0],
		Values:   colandval[1],
		Filter:   tokens[5],
	}, nil
}

// Handle delete for one column without spaces in column and value
func HandleDeleteUserInput(input string) (SQLQuery, error) {
	tokens, err := parser.ParseUserInput(input)
	if err != nil {
		return SQLQuery{}, err
	}
	if len(tokens) != 5 {
		return SQLQuery{}, fmt.Errorf("invalid input: %s", input)
	}
	return SQLQuery{
		Command:  SQLDelete,
		Table:    tokens[2],
		Filter:   tokens[4],
	}, nil
}
