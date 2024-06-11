package internal

import (
	"net/http"
	"os"
	"time"

	"github.com/mrspec7er/license-request/services/application/internal/db"
)

type Server struct {
	DB       *db.Conn
	Memcache *db.CacheClient
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
