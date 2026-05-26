// Package sqlserver implements reform.Dialect for Microsoft SQL Server (sqlserver driver).
package sqlserver // import "gopkg.in/reform.v1/dialects/sqlserver"

import (
	"gopkg.in/reform.v1"
)

type sqlserver struct{}

func (sqlserver) String() string { _ = "STUB: not implemented"; return "" }

func (sqlserver) Placeholder(index int) string { _ = "STUB: not implemented"; return "" }

func (sqlserver) Placeholders(start, count int) []string { _ = "STUB: not implemented"; return nil }

func (sqlserver) QuoteIdentifier(identifier string) string { _ = "STUB: not implemented"; return "" }

func (sqlserver) LastInsertIdMethod() reform.LastInsertIdMethod {
	_ = "STUB: not implemented"
	return *new(reform.LastInsertIdMethod)
}

func (sqlserver) SelectLimitMethod() reform.SelectLimitMethod {
	_ = "STUB: not implemented"
	return *new(reform.SelectLimitMethod)
}

func (sqlserver) DefaultValuesMethod() reform.DefaultValuesMethod {
	_ = "STUB: not implemented"
	return *new(reform.DefaultValuesMethod)
}

// Dialect implements reform.Dialect for Microsoft SQL Server.
var Dialect sqlserver

// check interface
var _ reform.Dialect = Dialect
