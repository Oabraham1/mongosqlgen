package mongo

import "fmt"

type Command string

const (
	MongoFind   Command = "find"
	MongoInsert Command = "insert"
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
