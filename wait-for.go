package main

import (
	"flag"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", 5432, "Resource port. Examples: 5432 for PostgreSQL, 3306 for MySQL, 5672 for RabbitMQ or 27017 for MongoDB")
	resourceType := flag.String("type", "postgres", "Resource type. Examples: postgres, mysql, rabbitmq, mongodb")
	user := flag.String("user", "postgres", "user")
	password := flag.String("password", "postgres", "password")
	name := flag.String("name", "postgres", "Resource name. It is the database name for PostgreSQL, MySQL or MongoDB and Virtual Host for RabbitMQ")
	maxAttempts := flag.Int("maxAttempts", 3, "Maximum attempts")
	seconds := flag.Int("seconds", 10, "Seconds to wait for Resource")

	flag.Parse()

	connected, err := connectTo(*resourceType, *host, *port, *user, *password, *name, *maxAttempts, *seconds)
	if err != nil {
		panic(err)
	}
	if !connected {
		os.Exit(1)
	}

}
