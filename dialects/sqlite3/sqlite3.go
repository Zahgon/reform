// Package sqlite3 implements reform.Dialect for SQLite3.
package sqlite3 // import "gopkg.in/reform.v1/dialects/sqlite3"

import (
	"gopkg.in/reform.v1"
)

type sqlite3 struct{}

func (sqlite3) String() string { _ = "STUB: not implemented"; return "" }

func (sqlite3) Placeholder(index int) string { _ = "STUB: not implemented"; return "" }

func (sqlite3) Placeholders(start, count int) []string { _ = "STUB: not implemented"; return nil }

func (sqlite3) QuoteIdentifier(identifier string) string { _ = "STUB: not implemented"; return "" }

func (sqlite3) LastInsertIdMethod() reform.LastInsertIdMethod {
	_ = "STUB: not implemented"
	return *new(reform.LastInsertIdMethod)
}

func (sqlite3) SelectLimitMethod() reform.SelectLimitMethod {
	_ = "STUB: not implemented"
	return *new(reform.SelectLimitMethod)
}

func (sqlite3) DefaultValuesMethod() reform.DefaultValuesMethod {
	_ = "STUB: not implemented"
	return *new(reform.DefaultValuesMethod)
}

// Dialect implements reform.Dialect for SQLite3.
var Dialect sqlite3

// check interface
var _ reform.Dialect = Dialect
