package internal

import (
	"net/http"
	"os"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Server struct {
	DB  *gorm.DB
	Hub *amqp091.Connection
}

func NewServer(s Server) *http.Server {
	go s.RegisterConsumersRoutes()

	server := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
