package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Connection Interface
type Connection interface {
	connect(host string, port int, user string, password string, name string) (bool, error)
}

func connectTo(resourceType string, host string, port int, user string, password string, name string, maxAttempts int, seconds int) (bool, error) {
	t, err := getResourceByType(resourceType)
	if err != nil {
		return false, err
	}
	connected := false
	attempt := 1
	for attempt <= maxAttempts && !connected {
		connected, err = t.connect(host, port, user, password, name)
		if err != nil {
			log.Fatal(err.Error())
		}
		if !connected {
			log.Printf("Attempt %d. Waiting for %d seconds...", attempt, seconds)
			time.Sleep(time.Duration(seconds) * time.Second)
		} else {
			log.Println("Connected to the resource")
		}
		attempt++
	}
	return connected, nil
}

func getResourceByType(resourceType string) (Connection, error) {
	if resourceType == "" {
		return nil, errors.New("resource type was not set")
	}
	if resourceType == "postgres" {
		return PostgreSQLConnection{}, nil
	} else if resourceType == "mysql" {
		return MySQLConnection{}, nil
	} else if resourceType == "rabbitmq" {
		return RabbitMQConnection{}, nil
	} else if resourceType == "mongodb" {
		return MongoDBConnection{}, nil
	}
	return nil, fmt.Errorf("invalid resource type: %s", resourceType)
}
