package src

import (
	"context"
	"encoding/json"
	"os"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/form/internal/hub"
	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct{}

func (p Publisher) Publish(queue string, body []byte, uid string) error {
	ch := hub.StartConnection()
	defer ch.Close()

	payload := amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
		Headers: amqp091.Table{
			"uid": uid,
		},
	}

	err := ch.PublishWithContext(context.Background(), os.Getenv("EXCHANGE_NAME"), queue, false, false, payload)
	if err != nil {
		return err
	}

	return nil
}

func (p Publisher) Create(form dto.Form, uid string) error {
	data, err := json.Marshal(form)
	if err != nil {
		return err
	}

	exName := os.Getenv("EXCHANGE_NAME")

	err = p.Publish(exName+".create", data, uid)
	if err != nil {
		return err
	}

	return nil
}
