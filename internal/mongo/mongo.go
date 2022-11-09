package mongo

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