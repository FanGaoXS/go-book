package main

import "strings"

type Database struct {
	users map[string]string
}

func NewDatabase() *Database {
	users := map[string]string{
		"admin": "admin",
	}
	return &Database{
		users: users,
	}
}

func (d *Database) Add(username string, password string) bool {
	if _, ok := d.users[username]; !ok {
		return false
	}

	d.users[username] = password
	return true
}

func (d Database) Get(username string, password string) bool {
	if pw, ok := d.users[username]; ok {
		return strings.EqualFold(password, pw)
	}

	return false
}

func (d Database) Len() int {
	return len(d.users)
}
