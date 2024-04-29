package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrspec7er/license-request/services/form/internal/src"
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func (s Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello There!"})
	})

	router.Route("/forms", src.ControllerModule(s.DB))

	return router
}

func (s Server) RegisterConsumersRoutes() {
	wg := sync.WaitGroup{}

	consumersRoutesConfig(s.Hub, &wg, s.DB)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Println("Interrupt signal received. Closing consumers...")
	wg.Wait()
	fmt.Println("All consumers closed.")
}

func consumersRoutesConfig(conn *amqp091.Connection, wg *sync.WaitGroup, DB *gorm.DB) {
	src.ConsumerModule(conn, wg, DB)
}
