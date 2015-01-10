package datastore

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

var (
	ConnectOnce sync.Once // Once mutex
	db          *sqlx.DB
)

// Established a database connection
func Connect() (*sqlx.DB, error) {
	var err error
	ConnectOnce.Do(func() {
		db, err = sqlx.Connect("mysql", "mathub:mathub@(d.docker:3306)/mathub")
	})

	return db, err
}
