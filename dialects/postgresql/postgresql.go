// Package postgresql implements reform.Dialect for PostgreSQL.
package postgresql // import "gopkg.in/reform.v1/dialects/postgresql"

import (
	"gopkg.in/reform.v1"
)

type postgresql struct{}

func (postgresql) String() string { _ = "STUB: not implemented"; return "" }

func (postgresql) Placeholder(index int) string { _ = "STUB: not implemented"; return "" }

func (postgresql) Placeholders(start, count int) []string { _ = "STUB: not implemented"; return nil }

func (postgresql) QuoteIdentifier(identifier string) string { _ = "STUB: not implemented"; return "" }

func (postgresql) LastInsertIdMethod() reform.LastInsertIdMethod {
	_ = "STUB: not implemented"
	return *new(reform.LastInsertIdMethod)
}

func (postgresql) SelectLimitMethod() reform.SelectLimitMethod {
	_ = "STUB: not implemented"
	return *new(reform.SelectLimitMethod)
}

func (postgresql) DefaultValuesMethod() reform.DefaultValuesMethod {
	_ = "STUB: not implemented"
	return *new(reform.DefaultValuesMethod)
}

// Dialect implements reform.Dialect for PostgreSQL.
var Dialect postgresql

// check interface
var _ reform.Dialect = Dialect
