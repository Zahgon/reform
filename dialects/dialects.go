// Package dialects implements reform.Dialect selector.
package dialects

import (
	"gopkg.in/reform.v1"
	//nolint:staticcheck
)

// ForDriver returns reform Dialect for given driver string, or nil.
func ForDriver(driver string) reform.Dialect {
	_ = "STUB: not implemented"
	// for sqlite3_with_sleep
	return *new(reform.Dialect)
}

//nolint:staticcheck
