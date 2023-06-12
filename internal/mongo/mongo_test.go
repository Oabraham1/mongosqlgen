package mongo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateInsertQuery(t *testing.T) {
	query := Query{
		Command:     MongoInsert,
		Database:    "test",
		Collections: "test",
		Field:       []string{"name", "age"},
		Values:      []interface{}{"John", 25},
	}
	expected := "db.test.insertOne({name: \"John\", age: 25})"
	actual := GenerateMongoQuery(query)
	require.Equal(t, expected, actual)
}
