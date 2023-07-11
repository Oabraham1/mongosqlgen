package generator

import (
	"fmt"

	"github.com/oabraham1/mongosqlgen/internal/converter"
	"github.com/oabraham1/mongosqlgen/internal/mongo"
	"github.com/oabraham1/mongosqlgen/internal/sql"
)

// GenerateMongoQueryFromSQLQuery generates a MongoDB query from a SQL query
func GenerateMongoQueryFromSQLQuery(input string) (string, error) {
	// Parse the input into a SQL Query
	sqlQuery, err := sql.ConvertUserInputToSQLQuery(input)
	if err != nil {
		return "", err
	}

	fmt.Printf("SQL Query: %+v\n", sqlQuery)

	// Convert the SQL Query into a Mongo Query
	mongoQuery, err := converter.ConvertSQLQueryToMongoQuery(sqlQuery)
	if err != nil {
		return "", err
	}

	fmt.Printf("Mongo Query: %+v\n", mongoQuery)

	// Generate the Mongo Query
	mongoQueryStr := mongo.GenerateMongoQuery(mongoQuery)

	return mongoQueryStr, nil
}
