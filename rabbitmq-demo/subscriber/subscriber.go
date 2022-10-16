package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	c "rabbitmq-demo/connection"
	"strings"
)

func main() {
	channel, conn := c.CreateConnection()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("insert queue name: ")
	queueName, err := reader.ReadString('\n')
	queueName = strings.Replace(queueName, "\n", "", -1)
	c.FailOnError(err, "fail on reading string")

	c.QueueSetUp(queueName, channel)

	deliveries, err := channel.Consume(
		queueName,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args)
	)
	c.FailOnError(err, "error on consuming")

	forever := make(chan struct{})

	go func() {
		for d := range deliveries {
			log.Printf("%+v", d.Body)
		}
	}()

	<-forever
	conn.Close()
}
