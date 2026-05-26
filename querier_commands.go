package reform

func filteredColumnsAndValues(str Struct, columnsIn []string, isUpdate bool) (columns []string, values []interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// select columns from set and collect values

// make error for extra columns

// TODO make exported type for that error

func (q *Querier) insert(str Struct, columns []string, values []interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// make query

// TODO optimize to avoid using reflection
// https://github.com/go-reform/reform/issues/269
// record.SetPK(id)

func (q *Querier) beforeInsert(str Struct) error { _ = "STUB: not implemented"; return nil }

// Insert inserts a struct into SQL database table.
// If str implements BeforeInserter, it calls BeforeInsert() before doing so.
//
// It fills record's primary key field.
func (q *Querier) Insert(str Struct) error { _ = "STUB: not implemented"; return nil }

// cut primary key

// InsertColumns inserts a struct into SQL database table with specified columns.
// Other columns are omitted from generated INSERT statement.
// If str implements BeforeInserter, it calls BeforeInsert() before doing so.
//
// It fills record's primary key field.
func (q *Querier) InsertColumns(str Struct, columns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// InsertMulti inserts several structs into SQL database table with single query.
// If they implement BeforeInserter, it calls BeforeInsert() before doing so.
//
// All structs should belong to the same view/table.
// All records should either have or not have primary key set.
// It doesn't fill primary key fields.
// Given all these limitations, most users should use Querier.Insert in a loop, not this method.
func (q *Querier) InsertMulti(structs ...Struct) error { _ = "STUB: not implemented"; return nil }

// check that view is the same

// check if all PK are present or all are absent

// cut last ", "

func (q *Querier) update(str Struct, columns []string, values []interface{}, tail string, args ...interface{}) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (q *Querier) beforeUpdate(str Struct) error { _ = "STUB: not implemented"; return nil }

// Update updates all columns of row specified by primary key in SQL database table with given record.
// If record implements BeforeUpdater, it calls BeforeUpdate() before doing so.
//
// Method returns ErrNoRows if no rows were updated.
// Method returns ErrNoPK if primary key is not set.
func (q *Querier) Update(record Record) error { _ = "STUB: not implemented"; return nil }

// cut primary key, make tail

// UpdateColumns updates specified columns of row specified by primary key in SQL database table with given record.
// Other columns are omitted from generated UPDATE statement.
// If record implements BeforeUpdater, it calls BeforeUpdate() before doing so.
//
// Method returns ErrNoRows if no rows were updated.
// Method returns ErrNoPK if primary key is not set.
func (q *Querier) UpdateColumns(record Record, columns ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO make exported type for that error

// make tail

// UpdateView updates specified columns of rows specified by tail and args in SQL database table with given struct,
// and returns a number of updated rows.
// Other columns are omitted from generated UPDATE statement.
// If struct implements BeforeUpdater, it calls BeforeUpdate() before doing so.
//
// Method never returns ErrNoRows.
func (q *Querier) UpdateView(str Struct, columns []string, tail string, args ...interface{}) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// TODO make exported type for that error

// Save saves record in SQL database table.
// If primary key is set, it first calls Update and checks if row was affected (matched).
// If primary key is absent or no row was affected, it calls Insert. This allows to call Save with Record
// with primary key set.
func (q *Querier) Save(record Record) error { _ = "STUB: not implemented"; return nil }

// Delete deletes record from SQL database table by primary key.
//
// Method returns ErrNoRows if no rows were deleted.
// Method returns ErrNoPK if primary key is not set.
func (q *Querier) Delete(record Record) error { _ = "STUB: not implemented"; return nil }

// DeleteFrom deletes rows from view with tail and args and returns a number of deleted rows.
//
// Method never returns ErrNoRows.
func (q *Querier) DeleteFrom(view View, tail string, args ...interface{}) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
