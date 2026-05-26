package reform

import (
	"context"
	"database/sql"
)

// DBInterface is a subset of *sql.DB used by reform.
// Can be used together with NewDBFromInterface for easier integration with existing code or for passing test doubles.
//
// It may grow and shrink over time to include only needed *sql.DB methods,
// and is excluded from SemVer compatibility guarantees.
type DBInterface interface {
	DBTXContext
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// Deprecated: do not use, it will be removed in v1.6.
	DBTX
	// Deprecated: do not use, it will be removed in v1.6.
	Begin() (*sql.Tx, error)
}

// check interface
var _ DBInterface = (*sql.DB)(nil)

// DB represents a connection to SQL database.
type DB struct {
	*Querier
	db DBInterface
}

// NewDB creates new DB object for given SQL database connection.
// Logger can be nil.
func NewDB(db *sql.DB, dialect Dialect, logger Logger) *DB { _ = "STUB: not implemented"; return nil }

// NewDBFromInterface creates new DB object for given DBInterface.
// Can be used for easier integration with existing code or for passing test doubles.
// Logger can be nil.
func NewDBFromInterface(db DBInterface, dialect Dialect, logger Logger) *DB {
	_ = "STUB: not implemented"
	return nil
}

// DBInterface returns DBInterface associated with a given DB object.
func (db *DB) DBInterface() DBInterface {
	_ = "STUB: not implemented"

	// Begin starts transaction with Querier's context and default options.
	return *new(DBInterface)
}

func (db *DB) Begin() (*TX, error) { _ = "STUB: not implemented"; return nil, nil }

// BeginTx starts transaction with given context and options (can be nil).
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*TX, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// InTransaction wraps function execution in transaction with Querier's context and default options,
// rolling back it in case of error or panic, committing otherwise.
func (db *DB) InTransaction(f func(t *TX) error) error { _ = "STUB: not implemented"; return nil }

// InTransactionContext wraps function execution in transaction with given context and options (can be nil),
// rolling back it in case of error or panic, committing otherwise.
func (db *DB) InTransactionContext(ctx context.Context, opts *sql.TxOptions, f func(t *TX) error) error {
	_ = "STUB: not implemented"
	return nil
}

// always return f() or Commit() error, not possible Rollback() error

// check interfaces
var (
	_ DBTX        = (*DB)(nil)
	_ DBTXContext = (*DB)(nil)
)
