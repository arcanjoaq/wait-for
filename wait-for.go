package main

import (
	"flag"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 5432, "Database port")
	dbtype := flag.String("dbtype", "postgres", "Database type")
	user := flag.String("user", "postgres", "Database user")
	password := flag.String("password", "postgres", "Database password")
	database := flag.String("database", "postgres", "Database name")

	maxAttempts := flag.Int("maxAttempts", 3, "Maximum attempts")
	seconds := flag.Int("seconds", 10, "Seconds to wait for Database")

	flag.Parse()

	connected, err := connectTo(*dbtype, *host, *port, *user, *password, *database, *maxAttempts, *seconds)
	if err != nil {
		panic(err)
	}
	if !connected {
		os.Exit(1)
	}

}
