package main

import (
	"database/sql"
)

//go:generate reform

type yesNo bool

func (yn *yesNo) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

// check interface
var _ sql.Scanner = (*yesNo)(nil)

//reform:information_schema.tables
type table struct {
	TableCatalog string `reform:"table_catalog"`
	TableSchema  string `reform:"table_schema"`
	TableName    string `reform:"table_name"`
	TableType    string `reform:"table_type"`
}

//reform:information_schema.columns
type column struct {
	TableCatalog string `reform:"table_catalog"`
	TableSchema  string `reform:"table_schema"`
	TableName    string `reform:"table_name"`
	Name         string `reform:"column_name"`
	IsNullable   yesNo  `reform:"is_nullable"`
	Type         string `reform:"data_type"`
}

//reform:information_schema.key_column_usage
type keyColumnUsage struct {
	ColumnName      string `reform:"column_name"`
	OrdinalPosition int    `reform:"ordinal_position"`
}

//reform:sqlite_master
type sqliteMaster struct {
	Name string `reform:"name"`
}

// TODO This "dummy" table name is ugly. We should do better.
// See https://github.com/go-reform/reform/issues/107.
//
//reform:dummy
type sqliteTableInfo struct {
	CID          int     `reform:"cid"`
	Name         string  `reform:"name"`
	Type         string  `reform:"type"`
	NotNull      bool    `reform:"notnull"`
	DefaultValue *string `reform:"dflt_value"`
	PK           int     `reform:"pk"`
}
