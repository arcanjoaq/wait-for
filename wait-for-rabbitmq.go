package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQConnection for RabbitMQ
type RabbitMQConnection struct{}

func (s RabbitMQConnection) connect(host string, port int, user string, password string, virtualHost string) bool {
	if port < 0 {
		port = 5672
	}
	connectionInfo := fmt.Sprintf("amqp://%s:%s@%s:%d%s", user, password, host, port, virtualHost)
	conn, err := amqp.Dial(connectionInfo)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return false
	}
	return true
}
