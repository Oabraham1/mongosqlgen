// Converts SQL query to MongoDB query.
package converter

import "fmt"
import "github.com/oabraham1/mongosqlgen/internal/sql"

type MongoCommand string

const (
	MongoFind   MongoCommand = "find"
	MongoInsert MongoCommand = "insert"
	MongoUpdate MongoCommand = "update"
	MongoDelete MongoCommand = "delete"
)

type MongoQuery struct {
	Command     MongoCommand
	Database    string
	Collections string
	Field       string
	Filter      string
	Values      interface{}
}

func ConvertSQLCommandToMongoCommand(command sql.SQLCommand) (MongoCommand, error) {
	switch command {
	case sql.SQLSelect:
		return MongoFind, nil
	case sql.SQLInsert:
		return MongoInsert, nil
	case sql.SQLUpdate:
		return MongoUpdate, nil
	case sql.SQLDelete:
		return MongoDelete, nil
	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}
}

func ConvertSQLQueryToMongoQuery(query sql.SQLQuery) (MongoQuery, error) {
	mongoCommand, err := ConvertSQLCommandToMongoCommand(query.Command)
	if err != nil {
		return MongoQuery{}, err
	}
	return MongoQuery{
		Command:     mongoCommand,
		Database:    query.Database,
		Collections: query.Table,
		Field:       query.Columns,
		Filter:      query.Filter,
		Values:      query.Values,
	}, nil

}
