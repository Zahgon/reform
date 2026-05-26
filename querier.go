package reform

import (
	"context"
	"database/sql"
	"time"
)

// Querier performs queries and commands.
type Querier struct {
	ctx     context.Context
	dbtxCtx DBTXContext
	tag     string
	Dialect
	Logger Logger
}

func newQuerier(ctx context.Context, dbtxCtx DBTXContext, tag string, dialect Dialect, logger Logger) *Querier {
	_ = "STUB: not implemented"
	return nil
}

func (q *Querier) clone() *Querier { _ = "STUB: not implemented"; return nil }

func (q *Querier) logBefore(query string, args []interface{}) { _ = "STUB: not implemented"; return }

func (q *Querier) logAfter(query string, args []interface{}, d time.Duration, err error) {
	_ = "STUB: not implemented"
	return
}

func (q *Querier) startQuery(command string) string { _ = "STUB: not implemented"; return "" }

// Tag returns Querier's tag. Default tag is empty.
func (q *Querier) Tag() string {
	_ = "STUB: not implemented"

	// WithTag returns a copy of Querier with set tag. Returned Querier is tied to the same DB or TX.
	// See Tagging section in documentation for details.
	return ""
}

func (q *Querier) WithTag(format string, args ...interface{}) *Querier {
	_ = "STUB: not implemented"
	return nil
}

// QualifiedView returns quoted qualified view name.
func (q *Querier) QualifiedView(view View) string { _ = "STUB: not implemented"; return "" }

// Context returns Querier's context. Default context is context.Background().
func (q *Querier) Context() context.Context {
	_ = "STUB: not implemented"

	// WithContext returns a copy of Querier with set context. Returned Querier is tied to the same DB or TX.
	// See Context section in documentation for details.
	return *new(context.Context)
}

func (q *Querier) WithContext(ctx context.Context) *Querier { _ = "STUB: not implemented"; return nil }

// QualifiedColumns returns a slice of quoted qualified column names for given view.
func (q *Querier) QualifiedColumns(view View) []string { _ = "STUB: not implemented"; return nil }

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (q *Querier) Exec(query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// ExecContext just calls q.WithContext(ctx).Exec(query, args...), and that form should be used instead.
// This method exists to satisfy various standard interfaces for advanced use-cases.
func (q *Querier) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (q *Querier) Query(query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryContext just calls q.WithContext(ctx).Query(query, args...), and that form should be used instead.
// This method exists to satisfy various standard interfaces for advanced use-cases.
func (q *Querier) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until Row's Scan method is called.
func (q *Querier) QueryRow(query string, args ...interface{}) *sql.Row {
	_ = "STUB: not implemented"
	return nil
}

// QueryRowContext just calls q.WithContext(ctx).QueryRow(query, args...), and that form should be used instead.
// This method exists to satisfy various standard interfaces for advanced use-cases.
func (q *Querier) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	_ = "STUB: not implemented"
	return nil
}

// check interfaces
var (
	_ DBTX        = (*Querier)(nil)
	_ DBTXContext = (*Querier)(nil)
)
