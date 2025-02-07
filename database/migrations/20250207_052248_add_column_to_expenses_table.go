package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToExpensesTable_20250207_052248 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToExpensesTable_20250207_052248{}
	m.Created = "20250207_052248"

	migration.Register("AddColumnToExpensesTable_20250207_052248", m)
}

// Run the migrations
func (m *AddColumnToExpensesTable_20250207_052248) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *AddColumnToExpensesTable_20250207_052248) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE expense_records add COLUMN receipt_image_path varchar(255) default null after category")
}
