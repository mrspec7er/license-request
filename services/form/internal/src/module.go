package src

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func ControllerModule(DB *gorm.DB) func(chi.Router) {
	c := FormController{
		Service: FormService{
			DB: DB,
		},
	}

	return func(r chi.Router) {
		r.Get("/", c.GetAll)
	}
}

func ConsumerModule(conn *amqp091.Connection, wg *sync.WaitGroup, DB *gorm.DB) {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// defer ch.Close()

	l := &FormConsumer{
		Service: FormService{
			DB: DB,
		},
	}

	exName := os.Getenv("EXCHANGE_NAME")
	serverID := os.Getenv("SERVER_ID")

	fmt.Println("LISTENER_NAME", exName+".create")

	wg.Add(1)
	go l.Create(ch, wg, exName+".create", exName+".create"+serverID)

}
