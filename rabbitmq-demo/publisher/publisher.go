package main

import (
	"bufio"
	"fmt"
	"os"
	c "rabbitmq-demo/connection"

	"github.com/streadway/amqp"
)

func main() {
	channel, connection := c.CreateConnection()
	defer connection.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("message: ")
		message, err := reader.ReadString('\n')
		c.FailOnError(err, "error reading message")

		err = channel.Publish("broadcast", "", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
		c.FailOnError(err, "error publishing the message")
	}
}
