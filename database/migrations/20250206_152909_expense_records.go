package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ExpenseRecords_20250206_152909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ExpenseRecords_20250206_152909{}
	m.Created = "20250206_152909"

	migration.Register("ExpenseRecords_20250206_152909", m)
}

// Run the migrations
func (m *ExpenseRecords_20250206_152909) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE expense_records(`expense_record_id` int(11) NOT NULL AUTO_INCREMENT,`category` int(11) DEFAULT NULL,`description` varchar(255) DEFAULT NULL,`currency_id` int(11) DEFAULT NULL,`amount` float(20,4) DEFAULT NULL,`expense_date` datetime DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) NOT NULL,`modified_by` int(11) NOT NULL,PRIMARY KEY (`expense_record_id`), FOREIGN KEY (currency_id) REFERENCES currencies(currency_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (created_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (modified_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (category) REFERENCES payment_categories(payment_category_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ExpenseRecords_20250206_152909) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `expense_records`")
}
