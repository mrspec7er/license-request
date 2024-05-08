package src

import (
	"context"
	"encoding/json"
	"os"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request/services/logger/internal/hub"
	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct{}

func (p Publisher) PublishLog(status int, uid string, reqPayload any, message string) error {
	ch := hub.StartConnection()
	defer ch.Close()

	payloadStringified, err := json.Marshal(reqPayload)
	if err != nil {
		return err
	}

	data := &dto.Logger{
		Status:  status,
		UID:     uid,
		Payload: string(payloadStringified),
		Message: message,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	payload := amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	}

	exName := os.Getenv("LOGGER_EXCHANGE")
	err = ch.PublishWithContext(context.Background(), exName, exName+".create", false, false, payload)
	if err != nil {
		return err
	}

	return nil
}
