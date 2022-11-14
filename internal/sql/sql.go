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
	Columns  []string
	Filter   string
	Values   []interface{}
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

func HandleSelectUserInput(input string) (Query, error) {
	var result Query

	result.Command = SQLSelect
	column, err := parser.FindBetween(input, "SELECT", "FROM")
	if err != nil {
		return Query{}, err
	}
	columns, err := parser.SplitInputByDelimiters(column, []string{" ", ","})
	if err != nil {
		return Query{}, err
	}
	result.Columns = columns

	has := parser.ContainsCommand(input, "WHERE")
	if has {
		tables, err := parser.FindBetween(input, "FROM", "WHERE")
		if err != nil {
			return Query{}, err
		}
		table, err := parser.SplitInputByDelimiters(tables, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Table = table[0]

		filters, err := parser.FindAfter(input, "WHERE")
		if err != nil {
			return Query{}, err
		}
		filter, err := parser.SplitInputByDelimiters(filters, []string{" ", ",", "'"})
		if err != nil {
			return Query{}, err
		}
		for _, f := range filter {
			result.Filter = result.Filter + f
		}
	} else {
		tables, err := parser.FindAfter(input, "FROM")
		if err != nil {
			return Query{}, err
		}
		table, err := parser.SplitInputByDelimiters(tables, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Table = table[0]
		result.Filter = ""
	}
	return result, nil
}

func HandleInsertUserInput(input string) (Query, error) {
	var result Query
	result.Command = SQLInsert
	if parser.CountOccurences(input, "(") == 1 {
		temp, err := parser.FindBetween(input, "INTO", "VALUES")
		if err != nil {
			return Query{}, err
		}
		table, err := parser.SplitInputByDelimiters(temp, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Table = table[0]
	} else {
		temp, err := parser.FindBetween(input, "INTO", "(")
		if err != nil {
			return Query{}, err
		}
		table, err := parser.SplitInputByDelimiters(temp, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Table = table[0]
	}

	if parser.CountOccurences(input, "(") == 2 {
		temp, err := parser.FindBetween(input, "(", ")")
		if err != nil {
			return Query{}, err
		}
		columns, err := parser.SplitInputByDelimiters(temp, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Columns = columns
	} else {
		result.Columns = []string{}
	}

	temp, err := parser.FindAfter(input, "VALUES")
	if err != nil {
		return Query{}, err
	}
	values, err := parser.SplitInputByDelimiters(temp, []string{" ", ",", "(", ")", "'"})
	if err != nil {
		return Query{}, err
	}
	for _, v := range values {
		result.Values = append(result.Values, v)
	}

	if len(result.Columns) > 0 {
		if len(result.Columns) != len(result.Values) {
			return Query{}, fmt.Errorf("number of columns and values do not match")
		}
	}
	return result, nil
}

func HandleUpdateUserInput(input string) (Query, error) {
	var result Query
	result.Command = SQLUpdate

	temp, err := parser.FindBetween(input, "UPDATE", "SET")
	if err != nil {
		return Query{}, err
	}
	table, err := parser.SplitInputByDelimiters(temp, []string{" ", ","})
	if err != nil {
		return Query{}, err
	}
	result.Table = table[0]

	if parser.ContainsCommand(input, "WHERE") {
		temp, err := parser.FindAfter(input, "WHERE")
		if err != nil {
			return Query{}, err
		}
		line, err := parser.SplitInputByDelimiters(temp, []string{" ", ",", "'"})
		if err != nil {
			return Query{}, err
		}
		for _, l := range line {
			result.Filter = result.Filter + l
		}
	} else {
		result.Filter = ""
	}

	var lines string
	if parser.ContainsCommand(input, "WHERE") {
		lines, err = parser.FindBetween(input, "SET", "WHERE")
		if err != nil {
			return Query{}, err
		}
	} else {
		lines, err = parser.FindAfter(input, "SET")
		if err != nil {
			return Query{}, err
		}
	}

	line, err := parser.SplitInputByDelimiters(lines, []string{","})
	if err != nil {
		return Query{}, err
	}
	for _, l := range line {
		column, err := parser.SplitInputByDelimiters(l, []string{"=", " ", "'"})
		if err != nil {
			return Query{}, err
		}
		result.Columns = append(result.Columns, column[0])
		result.Values = append(result.Values, column[1])
	}
	return result, nil
}

func HandleDeleteUserInput(input string) (Query, error) {
	var result Query

	if !parser.ContainsCommand(input, "DELETE") {
		return Query{}, fmt.Errorf("invalid command")
	}
	result.Command = SQLDelete

	if parser.ContainsCommand(input, "WHERE") {
		temp, err := parser.FindBetween(input, "FROM", "WHERE")
		if err != nil {
			return Query{}, err
		}
		table, err := parser.SplitInputByDelimiters(temp, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Table = table[0]

		temp, err = parser.FindAfter(input, "WHERE")
		if err != nil {
			return Query{}, err
		}
		line, err := parser.SplitInputByDelimiters(temp, []string{" ", ",", "'"})
		if err != nil {
			return Query{}, err
		}
		for _, l := range line {
			result.Filter = result.Filter + l
		}
	} else {
		temp, err := parser.FindAfter(input, "FROM")
		if err != nil {
			return Query{}, err
		}
		table, err := parser.SplitInputByDelimiters(temp, []string{" ", ","})
		if err != nil {
			return Query{}, err
		}
		result.Table = table[0]
		result.Filter = ""
	}

	return result, nil
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
