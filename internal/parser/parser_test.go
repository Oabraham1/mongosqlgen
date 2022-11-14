package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseUserInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:    "select",
			input:   "SELECT * FROM users",
			want:    []string{"SELECT", "*", "FROM", "users"},
			wantErr: false,
		},
		{
			name:    "insert",
			input:   "INSERT INTO users name VALUES 'Bob'",
			want:    []string{"INSERT", "INTO", "users", "name", "VALUES", "'Bob'"},
			wantErr: false,
		},
		{
			name:    "update",
			input:   "UPDATE users SET age = 30 WHERE name = 'Bob'",
			want:    []string{"UPDATE", "users", "SET", "age", "=", "30", "WHERE", "name", "=", "'Bob'"},
			wantErr: false,
		},
		{
			name:    "delete",
			input:   "DELETE FROM users WHERE name = 'Bob'",
			want:    []string{"DELETE", "FROM", "users", "WHERE", "name", "=", "'Bob'"},
			wantErr: false,
		},
		{
			name:    "invalid",
			input:   "SELECT",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserInput(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFindAfter(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		start   string
		want    string
		wantErr bool
	}{
		{
			name:    "select",
			input:   "SELECT * FROM users",
			start:   "SELECT",
			want:    " * FROM users",
			wantErr: false,
		},
		{
			name:    "insert",
			input:   "INSERT INTO users name VALUES 'Bob'",
			start:   "INTO",
			want:    " users name VALUES 'Bob'",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindAfter(tt.input, tt.start)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFindBetween(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		start   string
		end     string
		want    string
		wantErr bool
	}{
		{
			name:    "select",
			input:   "SELECT * FROM users",
			start:   "SELECT",
			end:     "FROM",
			want:    " * ",
			wantErr: false,
		},
		{
			name:    "insert",
			input:   "INSERT INTO users name VALUES 'Bob'",
			start:   "INTO",
			end:     "VALUES",
			want:    " users name ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBetween(tt.input, tt.start, tt.end)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestContainsCommand(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		subs    string
		want    bool
		wantErr bool
	}{
		{
			name:  "select",
			input: "SELECT * FROM users",
			subs:  "SELECT",
			want:  true,
		},
		{
			name:  "insert",
			input: "INSERT INTO users name VALUES 'Bob'",
			subs:  "INSERT",
			want:  true,
		},
		{
			name:  "update",
			input: "UPDATE users SET age = 30 WHERE name = 'Bob'",
			subs:  "VALUES",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsCommand(tt.input, tt.subs)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSplitInputByDelimiters(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:    "select",
			input:   "(name)",
			want:    []string{"name"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitInputByDelimiters(tt.input, []string{"(", ")", ",", " "})
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestCountOccurences(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		subs    string
		want    int
		wantErr bool
	}{
		{
			name:    "select",
			input:   "SELECT * FROM users",
			subs:    " ",
			want:    3,
			wantErr: false,
		},
		{
			name:    "insert",
			input:   "INSERT INTO users name VALUES 'Bob'",
			subs:    "'",
			want:    2,
			wantErr: false,
		},
		{
			name:    "update",
			input:   "UPDATE users",
			subs:    "(",
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOccurences(tt.input, tt.subs)
			require.Equal(t, tt.want, got)
		})
	}
}
