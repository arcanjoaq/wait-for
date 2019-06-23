package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	_ "github.com/lib/pq"
)

// PostgreSQLConnection Connector for PostgreSQL
type PostgreSQLConnection struct {}

func (s PostgreSQLConnection) connect(host string, port int, user string, password string, database string) bool {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, _ := sql.Open("postgres", dbinfo)
	err := db.Ping()
	if err != nil {
		cause := "Unknown"
		if strings.Contains(err.Error(), "connection refused") {
			cause = "Connection Refused"
		} else if strings.Contains(err.Error(), "connection reset by peer") {
			cause = "Connection Reset"
		}
		log.Printf("Error: %s", cause)
		return false
	} 
	return true
}
