package reform

import (
	"time"
)

// Inspect returns suitable for logging representation of a query argument.
func Inspect(arg interface{}, addType bool) string { _ = "STUB: not implemented"; return "" }

// Logger is responsible to log queries before and after their execution.
type Logger interface {
	// Before logs query before execution.
	Before(query string, args []interface{})

	// After logs query after execution.
	After(query string, args []interface{}, d time.Duration, err error)
}

// Printf is a (fmt.Printf|log.Printf|testing.T.Logf)-like function.
type Printf func(format string, args ...interface{})

// PrintfLogger is a simple query logger.
type PrintfLogger struct {
	LogTypes bool
	printf   Printf
}

// NewPrintfLogger creates a new simple query logger for any Printf-like function.
func NewPrintfLogger(printf Printf) *PrintfLogger { _ = "STUB: not implemented"; return nil }

// Before logs query before execution.
func (pl *PrintfLogger) Before(query string, args []interface{}) {
	_ = "STUB: not implemented"
	// fast path
	return
}

// After logs query after execution.
func (pl *PrintfLogger) After(query string, args []interface{}, d time.Duration, err error) {
	_ = "STUB: not implemented"
	// fast path
	return
}

// check interface
var _ Logger = (*PrintfLogger)(nil)
