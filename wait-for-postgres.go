package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgreSQLConnection for PostgreSQL
type PostgreSQLConnection struct{}

func (s PostgreSQLConnection) connect(host string, port int, user string, password string, database string) bool {
	if port < 0 {
		port = 5432
	}
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, _ := sql.Open("postgres", dbinfo)
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
