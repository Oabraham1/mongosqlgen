package sql

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseSQLCommand(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    Command
		wantErr bool
	}{
		{
			name:    "select",
			command: "SELECT",
			want:    SQLSelect,
			wantErr: false,
		},
		{
			name:    "insert",
			command: "INSERT",
			want:    SQLInsert,
			wantErr: false,
		},
		{
			name:    "update",
			command: "UPDATE",
			want:    SQLUpdate,
			wantErr: false,
		},
		{
			name:    "delete",
			command: "DELETE",
			want:    SQLDelete,
			wantErr: false,
		},
		{
			name:    "unknown",
			command: "UNKNOWN",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSQLCommand(tt.command)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetCommandFromUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Command
		wantErr bool
	}{
		{
			name:    "select",
			input:   "SELECT * FROM users",
			want:    SQLSelect,
			wantErr: false,
		},
		{
			name:    "insert",
			input:   "INSERT INTO users (name, age) VALUES ('Bob', 20)",
			want:    SQLInsert,
			wantErr: false,
		},
		{
			name:    "update",
			input:   "UPDATE users SET age = 21 WHERE name = 'Bob'",
			want:    SQLUpdate,
			wantErr: false,
		},
		{
			name:    "delete",
			input:   "DELETE FROM users WHERE name = 'Bob'",
			want:    SQLDelete,
			wantErr: false,
		},
		{
			name:    "unknown",
			input:   "UNKNOWN",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommandFromUserInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHandleSelectUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Query
		wantErr bool
	}{
		{
			name:    "select all",
			input:   "SELECT * FROM users",
			want:    Query{Command: SQLSelect, Database: "", Table: "users", Columns: []string{"*"}, Filter: ""},
			wantErr: false,
		},
		{
			name:    "select with filter",
			input:   "SELECT * FROM users WHERE name = 'Bob'",
			want:    Query{Command: SQLSelect, Database: "", Table: "users", Columns: []string{"*"}, Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "select with one column",
			input:   "SELECT name FROM users",
			want:    Query{Command: SQLSelect, Database: "", Table: "users", Columns: []string{"name"}, Filter: ""},
			wantErr: false,
		},
		{
			name:    "select with columns",
			input:   "SELECT name, age FROM users",
			want:    Query{Command: SQLSelect, Database: "", Table: "users", Columns: []string{"name", "age"}, Filter: ""},
			wantErr: false,
		},
		{
			name:    "unknown",
			input:   "UNKNOWN",
			want:    Query{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleSelectUserInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHandleInsertUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Query
		wantErr bool
	}{
		{
			name:    "insert one column multiple values",
			input:   "INSERT INTO users (name) VALUES ('Bob')",
			want:    Query{Command: SQLInsert, Database: "", Table: "users", Columns: []string{"name"}, Values: []interface{}{"Bob"}},
			wantErr: false,
		},
		{
			name:    "insert two columns multiple values",
			input:   "INSERT INTO users (name, age) VALUES ('Bob', 20)",
			want:    Query{Command: SQLInsert, Database: "", Table: "users", Columns: []string{"name", "age"}, Values: []interface{}{"Bob", "20"}},
			wantErr: false,
		},
		{
			name:    "insert all columns multiple values",
			input:   "INSERT INTO users VALUES ('Bob', 20)",
			want:    Query{Command: SQLInsert, Database: "", Table: "users", Columns: []string{}, Values: []interface{}{"Bob", "20"}},
			wantErr: false,
		},
		{
			name:    "insert mismatched columns and values",
			input:   "INSERT INTO users (name, age) VALUES ('Bob')",
			want:    Query{},
			wantErr: true,
		},
		{
			name:    "unknown",
			input:   "UNKNOWN",
			want:    Query{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleInsertUserInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHandleUpdateUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Query
		wantErr bool
	}{
		{
			name:    "update one column",
			input:   "UPDATE users SET age = 21 WHERE name = 'Bob'",
			want:    Query{Command: SQLUpdate, Database: "", Table: "users", Columns: []string{"age"}, Values: []interface{}{"21"}, Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "update two columns",
			input:   "UPDATE users SET age = 21, name = 'Bob' WHERE name = 'Bob'",
			want:    Query{Command: SQLUpdate, Database: "", Table: "users", Columns: []string{"age", "name"}, Values: []interface{}{"21", "Bob"}, Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "update with no filter",
			input:   "UPDATE users SET age = 21, name = 'Bob'",
			want:    Query{Command: SQLUpdate, Database: "", Table: "users", Columns: []string{"age", "name"}, Values: []interface{}{"21", "Bob"}, Filter: ""},
			wantErr: false,
		},
		{
			name:    "unknown",
			input:   "UNKNOWN",
			want:    Query{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleUpdateUserInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHandleDeleteUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Query
		wantErr bool
	}{
		{
			name:    "delete",
			input:   "DELETE FROM users WHERE name = 'Bob'",
			want:    Query{Command: SQLDelete, Database: "", Table: "users", Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "delete with no filter",
			input:   "DELETE FROM users",
			want:    Query{Command: SQLDelete, Database: "", Table: "users", Filter: ""},
			wantErr: false,
		},
		{
			name:    "unknown",
			input:   "UNKNOWN",
			want:    Query{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HandleDeleteUserInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCovertUserInputToSQLQuery(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Query
		wantErr bool
	}{
		{
			name:    "select",
			input:   "SELECT * FROM users WHERE name = 'Bob'",
			want:    Query{Command: SQLSelect, Database: "", Table: "users", Columns: []string{"*"}, Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "insert",
			input:   "INSERT INTO users (name, age) VALUES ('Bob', 20)",
			want:    Query{Command: SQLInsert, Database: "", Table: "users", Columns: []string{"name", "age"}, Values: []interface{}{"Bob", "20"}},
			wantErr: false,
		},
		{
			name:    "update",
			input:   "UPDATE users SET age = 21 WHERE name = 'Bob'",
			want:    Query{Command: SQLUpdate, Database: "", Table: "users", Columns: []string{"age"}, Values: []interface{}{"21"}, Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "delete",
			input:   "DELETE FROM users WHERE name = 'Bob'",
			want:    Query{Command: SQLDelete, Database: "", Table: "users", Filter: "name=Bob"},
			wantErr: false,
		},
		{
			name:    "unknown",
			input:   "UNKNOWN",
			want:    Query{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CovertUserInputToSQLQuery(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
