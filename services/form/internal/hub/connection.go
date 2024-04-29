package hub

import (
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func StartConnection() *amqp091.Connection {
	conn, err := amqp091.Dial(os.Getenv("MESSAGE_BROKER_URI"))
	if err != nil {
		panic(err)
	}

	return conn
}
