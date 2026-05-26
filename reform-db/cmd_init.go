package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/reform.v1"
	//nolint:staticcheck
)

var (
	initFlags = flag.NewFlagSet("init", flag.ExitOnError)
	gofmtF    = initFlags.Bool("gofmt", true, "Format with gofmt")
)

func init() {
	initFlags.Usage = func() {
		fmt.Fprintf(os.Stderr, "`init` generates Go model files for existing database schema.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [global flags] init [init flags] [directory]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Global flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nInit flags:\n")
		initFlags.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
It uses information_schema or similar RDBMS mechanism to inspect database
structure. For each table, it generates a single file with single struct type
definition with fields, types, and tags. Generated code then should be checked
and edited manually.
`)
	}
}

func gofmt(path string) { _ = "STUB: not implemented"; return }

type typeFunc func(sqlType string, nullable bool) (typ string, pack string, comment string)

// maybePointer returns *typ if nullable, typ otherwise.
func maybePointer(typ string, nullable bool) string { _ = "STUB: not implemented"; return "" }

// convertName converts snake_case name of table or column to CamelCase name of type or field.
// It also handles "_id" to "ID" conversion as a typical special case.
func convertName(sqlName string) string { _ = "STUB: not implemented"; return "" }

// getPrimaryKeyColumn returns single primary key column for given table, or nil.
func getPrimaryKeyColumn(db *reform.DB, catalog, schema, tableName string) *keyColumnUsage {
	_ = "STUB: not implemented"
	return nil
}

// MySQL doesn't have table_catalog in table_constraints

// get only the first row (with the maximum ordinal_position)

// initModelsInformationSchema returns structs from database with information_schema.
func initModelsInformationSchema(db *reform.DB, tablesTail string, typeFunc typeFunc) (structs []StructData) {
	_ = "STUB: not implemented"
	return nil
}

// cmdInit implements init command.
func cmdInit(db *reform.DB, dir string) { _ = "STUB: not implemented"; return }

// catalog is a currently selected database (reform-database, postgres, template0, etc.)
// schema is a PostgreSQL schema (public, pg_catalog, information_schema, etc.)

// catalog is always "def"
// schema is a database name (reform-database, information_schema, performance_schema, mysql, sys, etc.)

// SQLite is special

//nolint:staticcheck

// catalog is a currently selected database (reform-database, master, etc.)
// schema is MS SQL schema (dbo, guest, sys, information_schema, etc.)

// detect package name by importing package or from directory name
