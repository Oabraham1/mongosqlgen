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
			want:    Query{Command: SQLSelect, Database: "", Table: "users", Columns: "*", Filter: ""},
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
			name:    "insert",
			input:   "INSERT INTO users (name) VALUES (Bob)",
			want:    Query{Command: SQLInsert, Database: "", Table: "users", Columns: "name", Filter: "", Values: "Bob"},
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
			name:    "update",
			input:   "UPDATE users SET name=Bob WHERE name=Alice",
			want:    Query{Command: SQLUpdate, Database: "", Table: "users", Columns: "name", Filter: "name=Alice", Values: "Bob"},
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
			input:   "DELETE FROM users WHERE name=Bob",
			want:    Query{Command: SQLDelete, Database: "", Table: "users", Columns: "", Filter: "name=Bob"},
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
