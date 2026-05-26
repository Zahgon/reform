// Package test provides shared testing utilities.
package test

import (
	"sync"

	"gopkg.in/reform.v1"
	//nolint:staticcheck
)

//nolint:gochecknoglobals
var (
	sqlite3RegisterOnce sync.Once
	inspectOnce         sync.Once
)

// ConnectToTestDB returns open and prepared connection to test DB.
func ConnectToTestDB() *reform.DB { _ = "STUB: not implemented"; return nil }

// register custom function "sleep" for context tests

// Use single connection so various session-related variables work.
// For example: "PRAGMA foreign_keys" for SQLite3, "SET IDENTITY_INSERT" for MS SQL, etc.

// select dialect for driver

//nolint:staticcheck
