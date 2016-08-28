package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Clients_20160827_213836 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Clients_20160827_213836{}
	m.Created = "20160827_213836"
	migration.Register("Clients_20160827_213836", m)
}

// Run the migrations
func (m *Clients_20160827_213836) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE clients(id serial primary key,name TEXT NOT NULL,email TEXT NOT NULL)")
}

// Reverse the migrations
func (m *Clients_20160827_213836) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE clients")
}
