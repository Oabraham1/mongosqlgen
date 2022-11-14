package converter

import (
	"github.com/oabraham1/mongosqlgen/internal/mongo"
	"github.com/oabraham1/mongosqlgen/internal/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertSQLCommandToMongoCommand(t *testing.T) {
	tests := []struct {
		name    string
		command sql.Command
		want    mongo.Command
		wantErr bool
	}{
		{
			name:    "select",
			command: sql.SQLSelect,
			want:    mongo.MongoFind,
			wantErr: false,
		},
		{
			name:    "insert",
			command: sql.SQLInsert,
			want:    mongo.MongoInsert,
			wantErr: false,
		},
		{
			name:    "update",
			command: sql.SQLUpdate,
			want:    mongo.MongoUpdate,
			wantErr: false,
		},
		{
			name:    "delete",
			command: sql.SQLDelete,
			want:    mongo.MongoDelete,
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
			got, err := ConvertSQLCommandToMongoCommand(tt.command)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestConvertSQLQueryToMongoQuery(t *testing.T) {
	tests := []struct {
		name    string
		sql     sql.Query
		want    mongo.Query
		wantErr bool
	}{
		{
			name: "select",
			sql: sql.Query{
				Command:  sql.SQLSelect,
				Database: "test",
				Table:    "users",
				Columns:  []string{"name, age"},
			},
			want: mongo.Query{
				Command:     mongo.MongoFind,
				Database:    "test",
				Collections: "users",
				Field:       []string{"name, age"},
			},
			wantErr: false,
		},
		{
			name: "insert",
			sql: sql.Query{
				Command:  sql.SQLInsert,
				Database: "test",
				Table:    "users",
				Columns:  []string{"name"},
				Values:   []interface{}{"John"},
			},
			want: mongo.Query{
				Command:     mongo.MongoInsert,
				Database:    "test",
				Collections: "users",
				Field:       []string{"name"},
				Values:      []interface{}{"John"},
			},
			wantErr: false,
		},
		{
			name: "update",
			sql: sql.Query{
				Command:  sql.SQLUpdate,
				Database: "test",
				Table:    "users",
				Columns:  []string{"name"},
				Filter:   "id = 1",
				Values:   []interface{}{"John"},
			},
			want: mongo.Query{
				Command:     mongo.MongoUpdate,
				Database:    "test",
				Collections: "users",
				Field:       []string{"name"},
				Filter:      "id = 1",
				Values:      []interface{}{"John"},
			},
			wantErr: false,
		},
		{
			name: "delete",
			sql: sql.Query{
				Command:  sql.SQLDelete,
				Database: "test",
				Table:    "users",
				Filter:   "id = 1",
			},
			want: mongo.Query{
				Command:     mongo.MongoDelete,
				Database:    "test",
				Collections: "users",
				Filter:      "id = 1",
			},
			wantErr: false,
		},
		{
			name: "unknown",
			sql: sql.Query{
				Command:  "UNKNOWN",
				Database: "test",
				Table:    "users",
				Columns:  []string{"name"},
			},
			want:    mongo.Query{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertSQLQueryToMongoQuery(tt.sql)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, got)
		})
	}
}
