package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Module(DB *gorm.DB, memcache *redis.Client) func(chi.Router) {
	cs := Consumer{
		Service: Service{
			DB: DB,
		},
	}

	go cs.Load()

	return func(r chi.Router) {}
}
