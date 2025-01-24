/**
 * @author Jose Nidhin
 */
package pg

import (
	"context"
	"database/sql"
	"regexp"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/lftzzzzfeng/fasms/db"
)

var _ db.Execer = &Exec{}

var sqlLogRe = regexp.MustCompile(`\s+`)

// Exec is a DB function wrapper to execute single SQL. It provides an
// SQLExecer interfaces.
type Exec struct {
	db     *sqlx.DB
	logger *zap.Logger
}

// NewExec creates a new Exec.
func NewExec(sqlxDB *sqlx.DB, logger *zap.Logger) (*Exec, error) {
	exec := &Exec{
		db:     sqlxDB,
		logger: logger,
	}

	return exec, nil
}

// getExecer returns an appropriate implementation of Execer.
func (exec *Exec) getExecer(_ context.Context) db.Execer {
	return exec.db
}

// getLogger returns a logger with context.
func (exec *Exec) getLogger(_ context.Context) *zap.Logger {
	return exec.logger
}

// Refer sqlx GetContext for details.
func (exec *Exec) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	execer := exec.getExecer(ctx)
	logger := exec.getLogger(ctx)

	start := time.Now()
	defer func() {
		logger.Debug("SQL statement",
			zap.String("sql", sqlLogStr(query)),
			zap.Any("args", args),
			zap.Float64("duration", sqlDuration(start)))
	}()

	return execer.GetContext(ctx, dest, query, args...)
}

// Refer NamedExecContext for details.
func (exec *Exec) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	execer := exec.getExecer(ctx)
	logger := exec.getLogger(ctx)

	start := time.Now()
	defer func() {
		logger.Debug("SQL statement",
			zap.String("sql", sqlLogStr(query)),
			zap.Any("args", arg),
			zap.Float64("duration", sqlDuration(start)))
	}()

	return execer.NamedExecContext(ctx, query, arg)
}

// Refer QueryRowxContext for details.
func (exec *Exec) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	execer := exec.getExecer(ctx)
	logger := exec.getLogger(ctx)

	start := time.Now()
	defer func() {
		logger.Debug("SQL statement",
			zap.String("sql", sqlLogStr(query)),
			zap.Any("args", args),
			zap.Float64("duration", sqlDuration(start)))
	}()

	return execer.QueryRowxContext(ctx, query, args...)
}

// Refer QueryxContext for details.
func (exec *Exec) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	execer := exec.getExecer(ctx)
	logger := exec.getLogger(ctx)

	start := time.Now()
	defer func() {
		logger.Debug("SQL statement",
			zap.String("sql", sqlLogStr(query)),
			zap.Any("args", args),
			zap.Float64("duration", sqlDuration(start)))
	}()

	return execer.QueryxContext(ctx, query, args...)
}

// Refer SelectContext.
func (exec *Exec) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	execer := exec.getExecer(ctx)
	logger := exec.getLogger(ctx)

	start := time.Now()
	defer func() {
		logger.Debug("SQL statement",
			zap.String("sql", sqlLogStr(query)),
			zap.Any("args", args),
			zap.Float64("duration", sqlDuration(start)))
	}()

	return execer.SelectContext(ctx, dest, query, args...)
}

// sqlLogStr converts multiline formatted SQL statements into clean single line.
func sqlLogStr(s string) string {
	s = sqlLogRe.ReplaceAllString(s, " ")
	s = strings.Trim(s, " ")
	return s
}

// sqlDuration calculates the duration from start and returns the value as
// milliseconds.
func sqlDuration(start time.Time) float64 {
	d := time.Since(start)
	return float64(d) / float64(time.Millisecond)
}
