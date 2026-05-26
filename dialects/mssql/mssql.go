// Package mssql implements reform.Dialect for Microsoft SQL Server (mssql driver).
//
// Deprecated: Use sqlserver dialect instead. https://github.com/denisenkom/go-mssqldb#deprecated
package mssql // import "gopkg.in/reform.v1/dialects/mssql"

import (
	"gopkg.in/reform.v1"
)

type mssql struct{}

func (mssql) String() string { _ = "STUB: not implemented"; return "" }

func (mssql) Placeholder(index int) string { _ = "STUB: not implemented"; return "" }

func (mssql) Placeholders(start, count int) []string { _ = "STUB: not implemented"; return nil }

func (mssql) QuoteIdentifier(identifier string) string { _ = "STUB: not implemented"; return "" }

func (mssql) LastInsertIdMethod() reform.LastInsertIdMethod {
	_ = "STUB: not implemented"
	return *new(reform.LastInsertIdMethod)
}

func (mssql) SelectLimitMethod() reform.SelectLimitMethod {
	_ = "STUB: not implemented"
	return *new(reform.SelectLimitMethod)
}

func (mssql) DefaultValuesMethod() reform.DefaultValuesMethod {
	_ = "STUB: not implemented"
	return *new(reform.DefaultValuesMethod)
}

// Dialect implements reform.Dialect for Microsoft SQL Server.
//
// Deprecated: Use sqlserver.Dialect instead. https://github.com/denisenkom/go-mssqldb#deprecated
var Dialect mssql

// check interface
var _ reform.Dialect = Dialect
