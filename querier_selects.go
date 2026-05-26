package reform

import (
	"database/sql"
)

// NextRow scans next result row from rows to str. If str implements AfterFinder, it also calls AfterFind().
// It is caller's responsibility to call rows.Close().
//
// If there is no next result row, it returns ErrNoRows. It also may return rows.Err(), rows.Scan()
// and AfterFinder errors.
//
// See SelectRows example for idiomatic usage.
func (q *Querier) NextRow(str Struct, rows *sql.Rows) error { _ = "STUB: not implemented"; return nil }

// selectQuery returns full SELECT query for given view and tail.
func (q *Querier) selectQuery(view View, tail string, limit1 bool) string {
	_ = "STUB: not implemented"
	return ""
}

// SelectOneTo queries str's View with tail and args and scans first result to str.
// If str implements AfterFinder, it also calls AfterFind().
//
// If there are no rows in result, it returns ErrNoRows. It also may return QueryRow(), Scan()
// and AfterFinder errors.
func (q *Querier) SelectOneTo(str Struct, tail string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// SelectOneFrom queries view with tail and args and scans first result to new Struct str.
// If str implements AfterFinder, it also calls AfterFind().
//
// If there are no rows in result, it returns nil, ErrNoRows. It also may return QueryRow(), Scan()
// and AfterFinder errors.
func (q *Querier) SelectOneFrom(view View, tail string, args ...interface{}) (Struct, error) {
	_ = "STUB: not implemented"
	return *new(Struct), nil
}

// SelectRows queries view with tail and args and returns rows. They can then be iterated with NextRow().
// It is caller's responsibility to call rows.Close().
//
// In case of error rows will be nil. Error is never ErrNoRows.
//
// See example for idiomatic usage.
func (q *Querier) SelectRows(view View, tail string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectAllFrom queries view with tail and args and returns a slice of new Structs.
// If view's Struct implements AfterFinder, it also calls AfterFind().
//
// In case of query error slice will be nil. If error is encountered during iteration,
// partial result and error will be returned. Error is never ErrNoRows.
func (q *Querier) SelectAllFrom(view View, tail string, args ...interface{}) (structs []Struct, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// findTail returns a tail of SELECT query for given view, column and arg.
func (q *Querier) findTail(view string, column string, arg interface{}, limit1 bool) (tail string, needArg bool) {
	_ = "STUB: not implemented"
	return "", false
}

// FindOneTo queries str's View with column and arg and scans first result to str.
// If str implements AfterFinder, it also calls AfterFind().
//
// If there are no rows in result, it returns ErrNoRows. It also may return QueryRow(), Scan()
// and AfterFinder errors.
func (q *Querier) FindOneTo(str Struct, column string, arg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// FindOneFrom queries view with column and arg and scans first result to new Struct str.
// If str implements AfterFinder, it also calls AfterFind().
//
// If there are no rows in result, it returns nil, ErrNoRows. It also may return QueryRow(), Scan()
// and AfterFinder errors.
func (q *Querier) FindOneFrom(view View, column string, arg interface{}) (Struct, error) {
	_ = "STUB: not implemented"
	return *new(Struct), nil
}

// FindRows queries view with column and arg and returns rows. They can then be iterated with NextRow().
// It is caller's responsibility to call rows.Close().
//
// In case of error rows will be nil. Error is never ErrNoRows.
//
// See SelectRows example for idiomatic usage.
func (q *Querier) FindRows(view View, column string, arg interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindAllFrom queries view with column and args and returns a slice of new Structs.
// If view's Struct implements AfterFinder, it also calls AfterFind().
//
// In case of query error slice will be nil. If error is encountered during iteration,
// partial result and error will be returned. Error is never ErrNoRows.
func (q *Querier) FindAllFrom(view View, column string, args ...interface{}) ([]Struct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindByPrimaryKeyTo queries record's Table with primary key and scans first result to record.
// If record implements AfterFinder, it also calls AfterFind().
//
// If there are no rows in result, it returns ErrNoRows. It also may return QueryRow(), Scan()
// and AfterFinder errors.
func (q *Querier) FindByPrimaryKeyTo(record Record, pk interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// FindByPrimaryKeyFrom queries table with primary key and scans first result to new Record.
// If record implements AfterFinder, it also calls AfterFind().
//
// If there are no rows in result, it returns nil, ErrNoRows. It also may return QueryRow(), Scan()
// and AfterFinder errors.
func (q *Querier) FindByPrimaryKeyFrom(table Table, pk interface{}) (Record, error) {
	_ = "STUB: not implemented"
	return *new(Record), nil
}

// Reload is a shortcut for FindByPrimaryKeyTo for given record.
func (q *Querier) Reload(record Record) error { _ = "STUB: not implemented"; return nil }

// Count queries view with tail and args and returns a number (COUNT(*)) of matching rows.
func (q *Querier) Count(view View, tail string, args ...interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
