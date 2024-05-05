package src

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/application/internal/hub"
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		go c.Delete(ch, exName+".delete", exName+".delete"+serverID)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		go c.UpdateStatus(ch, exName+".status", exName+".status"+serverID)
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
		uid, ok := data.Headers["uid"].(string)
		if !ok {
			c.Hub.PublishLog(400, "", nil, "User ID undefine")
			continue
		}

		app := &dto.Application{}
		err := json.Unmarshal(data.Body, &app)
		if err != nil {
			c.Hub.PublishLog(400, uid, app, "Invalid data type")
			continue
		}

		status, err := c.Service.Create(app)
		if err != nil {
			c.Hub.PublishLog(status, uid, app, err.Error())
			continue
		}
		fmt.Println("APPS: ", *app)
		c.Hub.PublishLog(status, uid, app, "Create Application")
	}
}

func (c *Consumer) UpdateStatus(ch *amqp091.Channel, queue string, tag string) {
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

		app := &ChangeStatusInput{}
		err := json.Unmarshal(data.Body, &app)
		if err != nil {
			c.Hub.PublishLog(400, uid, app, "Invalid data type")
			continue
		}

		status, err := c.Service.ChangeStatus(app.Number, app.Status, app.Note, uid)
		if err != nil {
			c.Hub.PublishLog(status, uid, app, err.Error())
			continue
		}
		c.Hub.PublishLog(status, uid, app, "Update Status")
	}
}

func (c *Consumer) Delete(ch *amqp091.Channel, queue string, tag string) {

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

		app := &dto.Application{}
		err := json.Unmarshal(data.Body, &app)
		if err != nil {
			c.Hub.PublishLog(400, uid, app, "Invalid data type")
			continue
		}

		status, err := c.Service.Delete(app)
		if err != nil {
			c.Hub.PublishLog(status, uid, app, err.Error())
			continue
		}
		c.Hub.PublishLog(status, uid, app, "Delete Application")
	}
}
