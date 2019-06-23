package main

import (
	"fmt"
	"log"
	"time"
)

// DbConnection Interface for all Database Connections
type DbConnection interface {
	connect(host string, port int, user string, password string, database string) bool
}

func connectTo(dbtype string, host string, port int, user string, password string, database string, maxAttempts int, seconds int) (bool, error) {
	t, err := getDbConnectionByType(dbtype)
	if err != nil {
		return false, err
	}
	connected := false
	attempt := 1
	for attempt <= maxAttempts && !connected {
		connected = t.connect(host, port, user, password, database)
		if !connected {
			log.Printf("Attempt %d. Waiting for %d seconds...", attempt, seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
		} else {
			log.Println("Connected to the database")
		}
		attempt++
	}
	return connected, nil
}

func getDbConnectionByType(dbtype string) (DbConnection, error) {
	if dbtype == "postgres" {
		return PostgreSQLConnection{}, nil
	} else if dbtype == "mysql" {
		return MySQLConnection{}, nil
	}
	return nil, fmt.Errorf("Invalid dbtype: %s", dbtype)
}
