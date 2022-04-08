package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("rabbitMQ")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()
	msg, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Received Messages:%s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to our Instance")
	fmt.Print("[*] - waiting for messages")
	<-forever
}
