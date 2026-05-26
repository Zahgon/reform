// Package parse implements parsing of Go structs in files and runtime.
//
// This package, despite containing exported types, methods and functions,
// is an internal part of implementation of 'reform' command, also used by generated files,
// and not a part of public stable API.
package parse // import "gopkg.in/reform.v1/parse"

// FieldInfo represents information about struct field.
type FieldInfo struct {
	Name   string // field name as defined in source file, e.g. Name
	Type   string // field type as defined in source file, e.g. string; always present for primary key, may be absent otherwise
	Column string // SQL database column name from "reform:" struct field tag, e.g. name
}

// fieldInfoInSync returns true if FieldInfo fields that are set by both file and runtime parser are equal.
func fieldInfoInSync(fi1, fi2 *FieldInfo) bool { _ = "STUB: not implemented"; return false }

// GoString returns struct field information as Go code string.
func (fi *FieldInfo) GoString() string { _ = "STUB: not implemented"; return "" }

// StructInfo represents information about struct.
type StructInfo struct {
	Type         string      // struct type as defined in source file, e.g. User
	SQLSchema    string      // SQL database schema name from magic "reform:" comment, e.g. public
	SQLName      string      // SQL database view or table name from magic "reform:" comment, e.g. users
	Fields       []FieldInfo // fields info
	PKFieldIndex int         // index of primary key field in Fields, -1 if none
}

// structInfoInSync returns true if FieldInfo fields that are set by both file and runtime parser are equal.
func structInfoInSync(si1, si2 *StructInfo) bool { _ = "STUB: not implemented"; return false }

// GoString returns struct information as Go code string.
func (s *StructInfo) GoString() string { _ = "STUB: not implemented"; return "" }

// Columns returns a new slice of column names.
func (s *StructInfo) Columns() []string { _ = "STUB: not implemented"; return nil }

// ColumnsGoString returns column names as Go code string.
func (s *StructInfo) ColumnsGoString() string { _ = "STUB: not implemented"; return "" }

// IsTable returns true if this object represent information for table, false for view.
func (s *StructInfo) IsTable() bool { _ = "STUB: not implemented"; return false }

// PKField returns a primary key field, panics for views.
func (s *StructInfo) PKField() FieldInfo { _ = "STUB: not implemented"; return *new(FieldInfo) }

// AssertUpToDate checks that given StructInfo matches given object.
// It is used during program initialization to check that generated files are up-to-date.
func AssertUpToDate(si *StructInfo, obj interface{}) { _ = "STUB: not implemented"; return }

// parseStructFieldTag is used by both file and runtime parsers
func parseStructFieldTag(tag string) (sqlName string, isPK bool) {
	_ = "STUB: not implemented"
	return "", false
}

// checkFields is used by both file and runtime parsers
func checkFields(res *StructInfo) error { _ = "STUB: not implemented"; return nil }
