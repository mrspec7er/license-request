package src

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/rabbitmq/amqp091-go"
)

type FormConsumer struct {
	Service FormService
}

func (l *FormConsumer) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, tag string) {
	defer wg.Done()

	messages, err := queue.ConsumeWithContext(context.Background(), queueName, tag, true, false, false, false, nil)

	if err != nil {
		fmt.Println("ERROR1 :", err)
	}

	for data := range messages {
		uid, ok := data.Headers["uid"].(string)
		if !ok {
			fmt.Println("ERROR2 :", err)
			continue
		}

		form := &dto.Form{}

		err := json.Unmarshal(data.Body, &form)
		if err != nil {
			fmt.Println("ERROR3 :", err)
			continue
		}

		fmt.Println("IDENTIFIER :", uid)
		fmt.Println("RESULT :", form)
	}
}
