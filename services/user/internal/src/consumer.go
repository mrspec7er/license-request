package src

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/user/internal/hub"
	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	Service Service
	Hub     Publisher
}

func (c Consumer) Load() {
	ch := hub.StartConnection()
	defer ch.Close()

	wg := sync.WaitGroup{}

	exName := os.Getenv("EXCHANGE_NAME")
	serverID := os.Getenv("SERVER_ID")

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Create(ch, exName+".create", exName+".create"+serverID)
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Println("Interrupt signal received. Closing consumers...")
	wg.Wait()
	fmt.Println("All consumers closed.")
}

func (c *Consumer) Create(ch *amqp091.Channel, queue string, tag string) {
	messages, err := ch.ConsumeWithContext(context.Background(), queue, tag, true, false, false, false, nil)
	if err != nil {
		c.Hub.PublishLog(500, "", nil, "Failed to get messages from queue")
	}

	for data := range messages {

		user := &dto.User{}
		err := json.Unmarshal(data.Body, &user)
		if err != nil {
			c.Hub.PublishLog(400, "Unauthorize", user, "Invalid data type")
			continue
		}

		status, err := c.Service.Create(user)
		if err != nil {
			c.Hub.PublishLog(status, "Unauthorize", user, err.Error())
			continue
		}
		c.Hub.PublishLog(status, user.ID, user, "Create User")
	}
}
