package main

import (
	"gopkg.in/reform.v1"
)

// goTypeSQLite3 converts given SQL type to Go type. https://www.sqlite.org/datatype3.html
func goTypeSQLite3(sqlType string, nullable bool) (typ string, pack string, comment string) {
	_ = "STUB: not implemented"
	// SQLite3 has quite unique dynamic type system with storage classes and type affinities.
	// In short:
	// * table columns don't have rigid types;
	// * value has storage class (null, integer, real, text, blob), which defines how value is stored on disk;
	// * table column has type affinity (text, numeric, integer, real, blob), which defines preferred storage class
	//   for values in this column;
	// * table column declared type in CREATE TABLE defines column affinity with a set of rules;
	// * table_info returns column declared type;
	// * we try to mirror SQLite's set of rules of defining column affinity from declared type to define Go type;
	// * we also extend this set of rules with some common SQL data types;
	// * it's not 100% accurate (because SQLite is dynamically typed), but it follows actual and best practices.
	return "", "", ""
}

// SQLite rules 1-4

// never a pointer

//nolint:misspell

// common SQL data types

// numeric, decimal, etc.

// bool, boolean, etc.

// date, datetime, timestamp, etc.

// logger.Fatalf("unhandled SQLite3 type %q", sqlType)
// never a pointer

// initModelsSQLite3 returns structs from SQLite3 database.
func initModelsSQLite3(db *reform.DB) (structs []StructData) { _ = "STUB: not implemented"; return nil }

// no placeholders for PRAGMA

// https://github.com/go-reform/reform/issues/180

// not PK
// nothing

// composite PK
