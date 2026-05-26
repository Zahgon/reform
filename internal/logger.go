package internal

import (
	"gopkg.in/reform.v1"
)

// Logger is our custom logger with Debugf method.
type Logger struct {
	printf reform.Printf
	debug  bool
}

// NewLogger creates a new logger.
func NewLogger(prefix string, debug bool) *Logger { _ = "STUB: not implemented"; return nil }

// Debugf prints message only when Logger debug flag is set to true.
func (l *Logger) Debugf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1) (or panic for debug logger).
func (l *Logger) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// panic instead of os.Exit(1) to see output (SQL queries, failed assertions, etc.) in tests
