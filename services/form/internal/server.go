package internal

import (
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func NewServer(s Server) *http.Server {

	server := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
