package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLConnection for MySQL
type MySQLConnection struct{}

func (s MySQLConnection) connect(host string, port int, user string, password string, database string) bool {
	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)
	db, _ := sql.Open("mysql", dbinfo)
	if db != nil {
	err := db.Ping()
		defer db.Close()
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return false
		}
	}
	return true
}
