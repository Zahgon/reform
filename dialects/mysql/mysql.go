// Package mysql implements reform.Dialect for MySQL.
package mysql // import "gopkg.in/reform.v1/dialects/mysql"

import (
	"gopkg.in/reform.v1"
)

type mysql struct{}

func (mysql) String() string { _ = "STUB: not implemented"; return "" }

func (mysql) Placeholder(index int) string { _ = "STUB: not implemented"; return "" }

func (mysql) Placeholders(start, count int) []string { _ = "STUB: not implemented"; return nil }

func (mysql) QuoteIdentifier(identifier string) string { _ = "STUB: not implemented"; return "" }

func (mysql) LastInsertIdMethod() reform.LastInsertIdMethod {
	_ = "STUB: not implemented"
	return *new(reform.LastInsertIdMethod)
}

func (mysql) SelectLimitMethod() reform.SelectLimitMethod {
	_ = "STUB: not implemented"
	return *new(reform.SelectLimitMethod)
}

func (mysql) DefaultValuesMethod() reform.DefaultValuesMethod {
	_ = "STUB: not implemented"
	return *

	// Dialect implements reform.Dialect for MySQL.
	new(reform.DefaultValuesMethod)
}

var Dialect mysql

// check interface
var _ reform.Dialect = Dialect
