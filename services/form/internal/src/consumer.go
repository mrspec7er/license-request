package src

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/form/internal/hub"
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
	go c.Create(ch, &wg, exName+".create", exName+".create"+serverID)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Println("Interrupt signal received. Closing consumers...")
	wg.Wait()
	fmt.Println("All consumers closed.")
}

func (c *Consumer) Create(ch *amqp091.Channel, wg *sync.WaitGroup, queue string, tag string) {
	defer wg.Done()

	messages, err := ch.ConsumeWithContext(context.Background(), queue, tag, true, false, false, false, nil)
	if err != nil {
		c.Hub.PublishLog(500, "", nil, "Failed to get messages from queue")
	}

	for data := range messages {
		uid, ok := data.Headers["uid"].(string)
		if !ok {
			c.Hub.PublishLog(400, "", nil, "User ID undefine")
			continue
		}

		form := &dto.Form{}
		err := json.Unmarshal(data.Body, &form)
		if err != nil {
			c.Hub.PublishLog(400, uid, form, "Invalid data type")
			continue
		}

		status, err := c.Service.Create(form)
		if err != nil {
			c.Hub.PublishLog(status, uid, form, err.Error())
			continue
		}
		c.Hub.PublishLog(status, uid, form, "Create Form")
	}
}
