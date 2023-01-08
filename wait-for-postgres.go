package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// PostgreSQLConnection for PostgreSQL
type PostgreSQLConnection struct{}

func (s PostgreSQLConnection) connect(host string, port int, user string, password string, database string) (bool, error) {
	if port < 0 {
		port = 5432
	}
	if user == "" {
		return false, errors.New("user name was not set")
	}
	if password == "" {
		return false, errors.New("password was not set")
	}
	if database == "" {
		return false, errors.New("database name was not set")
	}

	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, _ := sql.Open("postgres", connectionInfo)
	if db != nil {
		err := db.Ping()
		defer db.Close()
		if err != nil {
			log.Printf("Error: %s", err.Error())
			return false, nil
		}
	}
	return true, nil
}
