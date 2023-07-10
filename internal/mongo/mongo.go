package mongo

import (
	"fmt"

	"github.com/oabraham1/mongosqlgen/internal/parser"
)

type Command string

const (
	MongoFind   Command = "find"
	MongoInsert Command = "insert"
	MongoUpdate Command = "update"
	MongoDelete Command = "deleteOne"
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
		return generateFindQuery(query)
	case MongoInsert:
		return generateInsertQuery(query)
	case MongoUpdate:
		return generateUpdateQuery(query)
	case MongoDelete:
		return generateDeleteQuery(query)
	default:
		return ""
	}

}

func generateFindQuery(query Query) string {
	fieldsAndValues := ""
	for i, field := range query.Field {
		if field == "*" {
			if query.Filter == "" {
				return fmt.Sprintf("db.%s.%s({%s})", query.Collections, query.Command, fieldsAndValues)
			} else {
				filterArray, err := parser.SplitInputByDelimiters(query.Filter, []string{"=", ">", "<", "!=", ">=", "<="})
				if err != nil {
					return ""
				}
				for i, f := range filterArray {
					if i%2 == 0 {
						fieldsAndValues += fmt.Sprintf("%s: \"%s\"", f, filterArray[i+1])
					}
					if i != len(filterArray)-2 && len(filterArray) > 2 {
						fieldsAndValues += ", "
					}
				}
			}
		} else {
			if query.Filter == "" && len(query.Field) > 0 {
				// Add field: 1 to fieldsAndValues
				fieldsAndValues += fmt.Sprintf("%s: 1", field)
				if i != len(query.Field)-1 {
					fieldsAndValues += ", "
				}

				// Return the proper filtering command
				if i == len(query.Field)-1 {
					return fmt.Sprintf("db.%s.%s({}, {%s})", query.Collections, query.Command, fieldsAndValues)
				}
			} else if query.Filter != "" && len(query.Field) > 0 {
				filterArray, err := parser.SplitInputByDelimiters(query.Filter, []string{"=", ">", "<", "!=", ">=", "<="})
				if err != nil {
					return ""
				}
				var firstFilter string
				for j, f := range filterArray {
					if j%2 == 0 {
						firstFilter += fmt.Sprintf("%s: \"%s\"", f, filterArray[j+1])
					}
					if j != len(filterArray)-2 && len(filterArray) > 2 {
						firstFilter += ", "
					}
				}

				fieldsAndValues += fmt.Sprintf("%s: 1", field)
				if i != len(query.Field)-1 {
					fieldsAndValues += ", "
				}

				if i == len(query.Field)-1 {
					return fmt.Sprintf("db.%s.%s({%s}, {%s})", query.Collections, query.Command, firstFilter, fieldsAndValues)
				}
			}
		}
	}
	return fmt.Sprintf("db.%s.%s({%s})", query.Collections, query.Command, fieldsAndValues)
}

func generateInsertQuery(query Query) string {
	var fieldsAndValues string
	for i, field := range query.Field {
		switch query.Values[i].(type) {
		case string:
			fieldsAndValues += fmt.Sprintf("%s: \"%s\"", field, query.Values[i])
		case int:
			fieldsAndValues += fmt.Sprintf("%s: %d", field, query.Values[i])
		case float64:
			fieldsAndValues += fmt.Sprintf("%s: %f", field, query.Values[i])
		default:
			fieldsAndValues += fmt.Sprintf("%s: %s", field, query.Values[i])
		}
		if i != len(query.Field)-1 {
			fieldsAndValues += ", "
		}
	}
	return fmt.Sprintf("db.%s.%s({%s})", query.Collections, query.Command, fieldsAndValues)
}

func generateUpdateQuery(query Query) string {
	var fieldsAndValues string
	// Add filter
	if query.Filter != "" {
		filterArray, err := parser.SplitInputByDelimiters(query.Filter, []string{"=", ">", "<", "!=", ">=", "<="})
		if err != nil {
			return ""
		}
		for i, f := range filterArray {
			if i%2 == 0 {
				fieldsAndValues += fmt.Sprintf("%s: \"%s\"", f, filterArray[i+1])
			}
			if i != len(filterArray)-2 && len(filterArray) > 2 {
				fieldsAndValues += ", "
			}
		}
		fieldsAndValues += "}, {"
	}

	for i, field := range query.Field {
		switch query.Values[i].(type) {
		case string:
			fieldsAndValues += fmt.Sprintf("%s: \"%s\"", field, query.Values[i])
		case int:
			fieldsAndValues += fmt.Sprintf("%s: %d", field, query.Values[i])
		case float64:
			fieldsAndValues += fmt.Sprintf("%s: %f", field, query.Values[i])
		default:
			fieldsAndValues += fmt.Sprintf("%s: %s", field, query.Values[i])
		}
		if i != len(query.Field)-1 {
			fieldsAndValues += ", "
		}
	}
	return fmt.Sprintf("db.%s.%s({%s})", query.Collections, query.Command, fieldsAndValues)
}

func generateDeleteQuery(query Query) string {
	var fieldsAndValues string
	// Add filter
	if query.Filter != "" {
		filterArray, err := parser.SplitInputByDelimiters(query.Filter, []string{"=", ">", "<", "!=", ">=", "<="})
		if err != nil {
			return ""
		}
		for i, f := range filterArray {
			if i%2 == 0 {
				fieldsAndValues += fmt.Sprintf("%s: \"%s\"", f, filterArray[i+1])
			}
			if i != len(filterArray)-2 && len(filterArray) > 2 {
				fieldsAndValues += ", "
			}
		}
	}
	return fmt.Sprintf("db.%s.%s({%s})", query.Collections, query.Command, fieldsAndValues)
}
