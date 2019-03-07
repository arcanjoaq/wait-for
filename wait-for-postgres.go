package main

import (
	"flag"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func connectToPostgreSQL(host string, port int, user string, password string, 
	database string, maxAttempts int, seconds int) bool {
	connected := false
	attempt := 1
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	for attempt <= maxAttempts && !connected {
		db, _ := sql.Open("postgres", dbinfo)
		err := db.Ping()
		if err != nil {
			cause := "Unknown"
			if strings.Contains(err.Error(), "connection refused") {
				cause = "Connection Refused"
			} else if strings.Contains(err.Error(), "connection reset by peer") {
				cause = "Connection Reset"
			}
			log.Printf("Attempt %d. %s. Waiting for %d seconds...", attempt, cause, seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
		} else {
			log.Println("Connected to the database")
			connected = true
		}
		attempt++
	}
	return connected
}

func main() {
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 5432, "Database port")
	user := flag.String("user", "postgres", "Database user")
	password := flag.String("password", "postgres", "Database password")
	database := flag.String("database", "postgres", "Database name")

	maxAttempts := flag.Int("maxAttempts", 3, "Maximum attempts")
	seconds := flag.Int("seconds", 10, "Seconds to wait for Database")
	
	flag.Parse()

	connected := connectToPostgreSQL(*host, *port, *user, *password, *database, *maxAttempts, *seconds)
	if !connected {
		os.Exit(1)
	}
}
