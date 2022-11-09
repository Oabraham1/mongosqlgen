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