package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Execer is an interface that defines the functions that are available in
// both sqlx.DB and sqlx.Tx which can be used for executing SQL statements.
// For more details checkout sqlx documentation.
type Execer interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
