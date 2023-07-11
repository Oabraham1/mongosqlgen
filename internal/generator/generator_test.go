package generator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMongoQueryFromSQLQuery(t *testing.T) {
	// Test for a SELECT query
	input := "SELECT * FROM users"
	want := `db.users.find({})`
	got, err := GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test for a SELECT query with a WHERE clause
	input = "SELECT * FROM users WHERE firstName = 'John'"
	want = `db.users.find({firstName: "John"})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test select specific columns
	input = "SELECT firstName, lastName FROM users"
	want = `db.users.find({}, {firstName: 1, lastName: 1})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test select specific columns with a WHERE clause
	input = "SELECT firstName, lastName FROM users WHERE firstName = 'John'"
	want = `db.users.find({firstName: "John"}, {firstName: 1, lastName: 1})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test for an INSERT query
	input = "INSERT INTO users (firstName, lastName) VALUES ('John', 'Doe')"
	want = `db.users.insert({firstName: "John", lastName: "Doe"})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test for an UPDATE query with single column
	input = "UPDATE users SET firstName = 'John' WHERE lastName = 'Doe'"
	want = `db.users.update({lastName: "Doe"}, {$set: {firstName: "John"}})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test for an UPDATE query with multiple columns
	input = "UPDATE users SET firstName = 'John', lastName = 'Doe' WHERE lastName = 'Doe'"
	want = `db.users.update({lastName: "Doe"}, {$set: {firstName: "John", lastName: "Doe"}})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)

	// Test for a DELETE query for specific column
	input = "DELETE FROM users WHERE firstName = 'John'"
	want = `db.users.deleteOne({firstName: "John"})`
	got, err = GenerateMongoQueryFromSQLQuery(input)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
