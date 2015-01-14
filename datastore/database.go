package datastore

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
		db, err = sqlx.Open("mysql", GetEnv()+"?parseTime=true")
	})
	return db, err
}

func GetEnv() string {
	user := os.Getenv("MYSQL_USER")
	passwd := os.Getenv("MYSQL_PASSWORD")
	db := os.Getenv("MYSQL_DATABASE")
	address := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	port := os.Getenv("MYSQL_PORT_3306_TCP_PORT")
	return fmt.Sprintf("%s:%s@(%s:%s)/%s", user, passwd, address, port, db)
}

// Bullshit
func ImportDump(dumpPath string) {
	re := regexp.MustCompilePOSIX("^CREATE TABLE ((.*)(\n.+))*;")
	file, err := ioutil.ReadFile(dumpPath)
	if err != nil {
		log.Printf("Import: couldn't load dump, %v", err)
		return
	}

	tables := re.FindAllString(string(file), -1)
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			log.Printf("Import: couldn't import, %v", err)
		}
	}
}
