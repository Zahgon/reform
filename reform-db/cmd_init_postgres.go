package main

// goTypePostgres converts given SQL type to Go type. https://www.postgresql.org/docs/current/static/datatype.html
func goTypePostgres(sqlType string, nullable bool) (typ string, pack string, comment string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// never a pointer

// interval can't be mapped to time.Duration: https://github.com/lib/pq/issues/78

// logger.Fatalf("unhandled PostgreSQL type %q", sqlType)
// never a pointer
