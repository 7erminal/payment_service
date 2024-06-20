package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Payments_20240523_190138 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Payments_20240523_190138{}
	m.Created = "20240523_190138"

	migration.Register("Payments_20240523_190138", m)
}

// Run the migrations
func (m *Payments_20240523_190138) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE payments(`payment_id` int(11) NOT NULL AUTO_INCREMENT,`initiated_by` int(11) DEFAULT NULL,`sender` int(11) DEFAULT NULL,`reciever` int(11) DEFAULT NULL,`amount` float NOT NULL,`payment_method` int(11) DEFAULT NULL,`status` int(11) DEFAULT NULL,`payment_account` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`payment_id`), FOREIGN KEY (initiated_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (sender) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (reciever) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (payment_method) REFERENCES payment_methods(payment_method_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (status) REFERENCES status(status_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Payments_20240523_190138) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `payments`")
}
