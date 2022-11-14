package mongo

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
