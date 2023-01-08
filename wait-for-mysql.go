package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLConnection for MySQL
type MySQLConnection struct{}

func (s MySQLConnection) connect(host string, port int, user string, password string, database string) (bool, error) {
	if port < 0 {
		port = 3306
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

	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)

	db, _ := sql.Open("mysql", connectionInfo)
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
