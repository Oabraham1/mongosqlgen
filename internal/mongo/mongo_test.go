package mongo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSingleInsertQuery(t *testing.T) {
	query := Query{
		Command:     MongoInsert,
		Database:    "test",
		Collections: "test",
		Field:       []string{"name", "age"},
		Values:      []interface{}{"John", 25},
	}
	expected := "db.test.insert({name: \"John\", age: 25})"
	actual := GenerateMongoQuery(query)
	require.Equal(t, expected, actual)
}

func TestGenerateUpdateQuery(t *testing.T) {
	query := Query{
		Command:     MongoUpdate,
		Database:    "test",
		Collections: "test",
		Field:       []string{"name", "age"},
		Values:      []interface{}{"John", 25},
	}
	expected := "db.test.update({name: \"John\", age: 25})"
	actual := GenerateMongoQuery(query)
	require.Equal(t, expected, actual)
}

func TestGenerateDeleteQuery(t *testing.T) {
	query := Query{
		Command:     MongoDelete,
		Database:    "test",
		Collections: "test",
		Field:       []string{"name", "age"},
		Values:      []interface{}{"John", 25},
	}
	expected := "db.test.delete({name: \"John\", age: 25})"
	actual := GenerateMongoQuery(query)
	require.Equal(t, expected, actual)
}
