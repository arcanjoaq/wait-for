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
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, database)
	db, _ := sql.Open("postgres", dbinfo)
	err := db.Ping()
	defer db.Close()
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return false
	}
	return true
}
