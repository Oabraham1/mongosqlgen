package converter

import (
	"github.com/oabraham1/mongosqlgen/internal/sql"
	"github.com/oabraham1/mongosqlgen/internal/mongo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConvertSQLCommandToMongoCommand(t *testing.T) {
	tests := []struct {
		name    string
		command sql.SQLCommand
		want    mongo.MongoCommand
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
		sql     sql.SQLQuery
		want    mongo.MongoQuery
		wantErr bool
	}{
		{
			name: "select",
			sql: sql.SQLQuery{
				Command:  sql.SQLSelect,
				Database: "test",
				Table:    "users",
				Columns:  "name",
			},
			want: mongo.MongoQuery{
				Command:     mongo.MongoFind,
				Database:    "test",
				Collections: "users",
				Field:       "name",
			},
			wantErr: false,
		},
		{
			name: "insert",
			sql: sql.SQLQuery{
				Command:  sql.SQLInsert,
				Database: "test",
				Table:    "users",
				Columns:  "name",
				Values:   "John",
			},
			want: mongo.MongoQuery{
				Command:     mongo.MongoInsert,
				Database:    "test",
				Collections: "users",
				Field:       "name",
				Values:      "John",
			},
			wantErr: false,
		},
		{
			name: "update",
			sql: sql.SQLQuery{
				Command:  sql.SQLUpdate,
				Database: "test",
				Table:    "users",
				Columns:  "name",
				Filter:   "id = 1",
				Values:   "John",
			},
			want: mongo.MongoQuery{
				Command:     mongo.MongoUpdate,
				Database:    "test",
				Collections: "users",
				Field:       "name",
				Filter:      "id = 1",
				Values:      "John",
			},
			wantErr: false,
		},
		{
			name: "delete",
			sql: sql.SQLQuery{
				Command:  sql.SQLDelete,
				Database: "test",
				Table:    "users",
				Filter:   "id = 1",
			},
			want: mongo.MongoQuery{
				Command:     mongo.MongoDelete,
				Database:    "test",
				Collections: "users",
				Filter:      "id = 1",
			},
			wantErr: false,
		},
		{
			name: "unknown",
			sql: sql.SQLQuery{
				Command:  "UNKNOWN",
				Database: "test",
				Table:    "users",
				Columns:  "name",
			},
			want:    mongo.MongoQuery{},
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
