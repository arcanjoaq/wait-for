package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	resourceType := flag.String("type", "", "Resource type. Examples: postgres, mysql, rabbitmq, mongodb")
	host := flag.String("host", "localhost", "host")
	port := flag.Int("port", -1, "Resource port. Examples: 5432 for PostgreSQL, 3306 for MySQL, 5672 for RabbitMQ or 27017 for MongoDB")
	user := flag.String("user", "", "user")
	password := flag.String("password", "", "password")
	name := flag.String("name", "", "Resource name. It is the database name for PostgreSQL, MySQL or MongoDB and Virtual Host for RabbitMQ")
	maxAttempts := flag.Int("maxAttempts", 3, "Maximum attempts")
	seconds := flag.Int("seconds", 10, "Seconds to wait for Resource")

	flag.Parse()

	connected, err := connectTo(*resourceType, *host, *port, *user, *password, *name, *maxAttempts, *seconds)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}
	if !connected {
		os.Exit(1)
	}

}
