package converter

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    SQLCommand
		wantErr bool
	}{
		{
			name:    "select",
			command: "SELECT",
			want:    Select,
			wantErr: false,
		},
		{
			name:    "insert",
			command: "INSERT",
			want:    Insert,
			wantErr: false,
		},
		{
			name:    "update",
			command: "UPDATE",
			want:    Update,
			wantErr: false,
		},
		{
			name:    "delete",
			command: "DELETE",
			want:    Delete,
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
			got, err := ParseCommand(tt.command)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}