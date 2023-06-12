package mongo

import "fmt"

type Command string

const (
	MongoFind   Command = "find"
	MongoInsert Command = "insertOne"
	MongoUpdate Command = "update"
	MongoDelete Command = "delete"
)

type Query struct {
	Command     Command
	Database    string
	Collections string
	Field       []string
	Filter      string
	Values      []interface{}
}

func GenerateMongoQuery(query Query) string {
	switch query.Command {
	case MongoFind:
		return GenerateMongoFindQuery(query)
	case MongoInsert:
		return GenerateMongoInsertQuery(query)
	case MongoUpdate:
		return GenerateMongoUpdateQuery(query)
	case MongoDelete:
		return GenerateMongoDeleteQuery(query)
	default:
		return ""
	}
}

func GenerateMongoFindQuery(query Query) string {
	return fmt.Sprintf("db.%s.%s(%s)", query.Database, query.Collections, query.Filter)
}

func GenerateMongoInsertQuery(query Query) string {
	fieldsAndValues := ""
	for i, field := range query.Field {
		switch query.Values[i].(type) {
		case string:
			fieldsAndValues += fmt.Sprintf("%s: \"%s\"", field, query.Values[i])
		default:
			fieldsAndValues += fmt.Sprintf("%s: %v", field, query.Values[i])
		}
		if i != len(query.Field)-1 {
			fieldsAndValues += ", "
		}
	}
	return fmt.Sprintf("db.%s.%s({%s})", query.Collections, query.Command, fieldsAndValues)
}

func GenerateMongoUpdateQuery(query Query) string {
	fieldsAndValues := ""
	for i, field := range query.Field {
		fieldsAndValues += fmt.Sprintf("%s: %s", field, query.Values[i])
		if i != len(query.Field)-1 {
			fieldsAndValues += ", "
		}
	}
	return fmt.Sprintf("db.%s.%s.update({%s})", query.Collections, query.Command, fieldsAndValues)
}

func GenerateMongoDeleteQuery(query Query) string {
	return fmt.Sprintf("db.%s.%s(%s)", query.Collections, query.Command, query.Filter)
}
