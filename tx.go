package reform

import (
	"context"
	"database/sql"
)

// TXInterface is a subset of *sql.Tx used by reform.
// Can be used together with NewTXFromInterface for easier integration with existing code or for passing test doubles.
//
// It may grow and shrink over time to include only needed *sql.Tx methods,
// and is excluded from SemVer compatibility guarantees.
type TXInterface interface {
	DBTXContext
	Commit() error
	Rollback() error

	// Deprecated: do not use, it will be removed in v1.6.
	DBTX
}

// check interface
var _ TXInterface = (*sql.Tx)(nil)

// TX represents a SQL database transaction.
type TX struct {
	*Querier
	tx TXInterface
}

// NewTX creates new TX object for given SQL database transaction.
// Logger can be nil.
func NewTX(tx *sql.Tx, dialect Dialect, logger Logger) *TX { _ = "STUB: not implemented"; return nil }

// NewTXFromInterface creates new TX object for given TXInterface.
// Can be used for easier integration with existing code or for passing test doubles.
// Logger can be nil.
func NewTXFromInterface(tx TXInterface, dialect Dialect, logger Logger) *TX {
	_ = "STUB: not implemented"
	return nil
}

func newTX(ctx context.Context, tx TXInterface, dialect Dialect, logger Logger) *TX {
	_ = "STUB: not implemented"
	return nil
}

// Commit commits the transaction.
func (tx *TX) Commit() error { _ = "STUB: not implemented"; return nil }

// Rollback aborts the transaction.
func (tx *TX) Rollback() error { _ = "STUB: not implemented"; return nil }

// check interfaces
var (
	_ DBTX        = (*TX)(nil)
	_ DBTXContext = (*TX)(nil)
)
