package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type PaymentHistory_20240523_191957 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &PaymentHistory_20240523_191957{}
	m.Created = "20240523_191957"

	migration.Register("PaymentHistory_20240523_191957", m)
}

// Run the migrations
func (m *PaymentHistory_20240523_191957) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE payment_history(`payment_history_id` int(11) NOT NULL AUTO_INCREMENT,`payment_id` int(11) NOT NULL,`status` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`payment_history_id`), FOREIGN KEY (payment_id) REFERENCES payments(payment_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *PaymentHistory_20240523_191957) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `payment_history`")
}
