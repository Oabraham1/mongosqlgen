package converter

type SQLCommand string

const (
	Select SQLCommand = "SELECT"
	Insert SQLCommand = "INSERT"
	Update SQLCommand = "UPDATE"
	Delete SQLCommand = "DELETE"
)

type SQLConverter interface {
	Convert(command SQLCommand, table string, columns []string, values []interface{}) (string, error)
}

type SQLConverterFunc func(command SQLCommand, table string, columns []string, values []interface{}) (string, error)

func (f SQLConverterFunc) Convert(command SQLCommand, table string, columns []string, values []interface{}) (string, error) {
	return f(command, table, columns, values)
}
