package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Status_20240523_185401 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Status_20240523_185401{}
	m.Created = "20240523_185401"

	migration.Register("Status_20240523_185401", m)
}

// Run the migrations
func (m *Status_20240523_185401) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE status(`status_id` int(11) NOT NULL AUTO_INCREMENT,`status` varchar(128) NOT NULL,`status_code` varchar(128) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`status_id`))")
}

// Reverse the migrations
func (m *Status_20240523_185401) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `status`")
}
