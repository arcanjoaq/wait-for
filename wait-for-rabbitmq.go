package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQConnection for RabbitMQ
type RabbitMQConnection struct{}

func (s RabbitMQConnection) connect(host string, port int, user string, password string, virtualHost string) (bool, error) {
	if port < 0 {
		port = 5672
	}
	if virtualHost == "" {
		virtualHost = "/"
	}

	var connectionInfo string
	if user != "" && password != "" {
		connectionInfo = fmt.Sprintf("amqp://%s:%s@%s:%d%s", user, password, host, port, virtualHost)
	} else {
		connectionInfo = fmt.Sprintf("amqp://%s:%d%s", host, port, virtualHost)
	}

	conn, err := amqp.Dial(connectionInfo)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return false, nil
	}
	return true, nil
}
