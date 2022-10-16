package connection

import (
	"log"

	"github.com/streadway/amqp"
)

func CreateConnection() (*amqp.Channel, *amqp.Connection) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "error connecting to rabbitmq")

	channel, err := connection.Channel()
	FailOnError(err, "error on creating channel")

	err = channel.ExchangeDeclare(
		"broadcast",
		"fanout",
		true,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   //arguments
	)
	FailOnError(err, "failed to declare exchange")

	return channel, connection
}

func QueueSetUp(queueName string, channel *amqp.Channel) {
	queue, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	FailOnError(err, "can not declare exchange 'broadcast'")

	err = channel.QueueBind(
		queue.Name,  // queue name
		"",          // routing key
		"broadcast", // exchange
		false,
		nil,
	)
	FailOnError(err, "can not bind queue to exchange 'broadcast'")

}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
