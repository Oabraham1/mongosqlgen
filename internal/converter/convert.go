package converter

import (
	"fmt"

	"github.com/oabraham1/mongosqlgen/internal/mongo"
	"github.com/oabraham1/mongosqlgen/internal/sql"
)

// ConvertSQLCommandToMongoCommand converts a SQL command to a MongoDB command
func ConvertSQLCommandToMongoCommand(command sql.Command) (mongo.Command, error) {
	switch command {
	case sql.SQLSelect:
		return mongo.MongoFind, nil
	case sql.SQLInsert:
		return mongo.MongoInsert, nil
	case sql.SQLUpdate:
		return mongo.MongoUpdate, nil
	case sql.SQLDelete:
		return mongo.MongoDelete, nil
	default:
		return "", fmt.Errorf("unknown command: %s", command)
	}
}

// ConvertSQLQueryToMongoQuery converts a SQL query to a MongoDB query
func ConvertSQLQueryToMongoQuery(query sql.Query) (mongo.Query, error) {
	mongoCommand, err := ConvertSQLCommandToMongoCommand(query.Command)
	if err != nil {
		return mongo.Query{}, err
	}
	return mongo.Query{
		Command:     mongoCommand,
		Database:    query.Database,
		Collections: query.Table,
		Field:       query.Columns,
		Filter:      query.Filter,
		Values:      query.Values,
	}, nil
}
