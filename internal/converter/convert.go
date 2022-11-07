// Converts SQL query to MongoDB query.
package converter

import "fmt"

type SQLCommand string

const (
	Select SQLCommand = "SELECT"
	Insert SQLCommand = "INSERT"
	Update SQLCommand = "UPDATE"
	Delete SQLCommand = "DELETE"
)

type SQLConverter interface {
	Convert(command SQLCommand, table string, columns []string, values []interface{}) (string, error)
}

type SQLConverterFunc func(command SQLCommand, table string, columns []string, values []interface{}) (string, error)

func (f SQLConverterFunc) Convert(command SQLCommand, table string, columns []string, values []interface{}) (string, error) {
	return f(command, table, columns, values)
}

func ParseCommand(command string) (SQLCommand, error) {
	switch command {
	case "SELECT":
		return Select, nil
	case "INSERT":
		return Insert, nil
	case "UPDATE":
		return Update, nil
	case "DELETE":
		return Delete, nil
	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}
}
