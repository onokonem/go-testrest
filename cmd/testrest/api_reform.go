package main

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type accountTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *accountTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("account").
func (v *accountTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *accountTableType) Columns() []string {
	return []string{"id", "name", "amount", "ctime", "mtime"}
}

// NewStruct makes a new struct for that view or table.
func (v *accountTableType) NewStruct() reform.Struct {
	return new(Account)
}

// NewRecord makes a new record for that table.
func (v *accountTableType) NewRecord() reform.Record {
	return new(Account)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *accountTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// AccountTable represents account view or table in SQL database.
var AccountTable = &accountTableType{
	s: parse.StructInfo{Type: "Account", SQLSchema: "", SQLName: "account", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int64", Column: "id"}, {Name: "Name", PKType: "", Column: "name"}, {Name: "Amount", PKType: "", Column: "amount"}, {Name: "Ctime", PKType: "", Column: "ctime"}, {Name: "Mtime", PKType: "", Column: "mtime"}}, PKFieldIndex: 0},
	z: new(Account).Values(),
}

// String returns a string representation of this struct or record.
func (s Account) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "Amount: " + reform.Inspect(s.Amount, true)
	res[3] = "Ctime: " + reform.Inspect(s.Ctime, true)
	res[4] = "Mtime: " + reform.Inspect(s.Mtime, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Account) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.Amount,
		s.Ctime,
		s.Mtime,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Account) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.Amount,
		&s.Ctime,
		&s.Mtime,
	}
}

// View returns View object for that struct.
func (s *Account) View() reform.View {
	return AccountTable
}

// Table returns Table object for that record.
func (s *Account) Table() reform.Table {
	return AccountTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Account) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Account) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Account) HasPK() bool {
	return s.ID != AccountTable.z[AccountTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Account) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int64(i64)
	} else {
		s.ID = pk.(int64)
	}
}

// check interfaces
var (
	_ reform.View   = AccountTable
	_ reform.Struct = new(Account)
	_ reform.Table  = AccountTable
	_ reform.Record = new(Account)
	_ fmt.Stringer  = new(Account)
)

func init() {
	parse.AssertUpToDate(&AccountTable.s, new(Account))
}
