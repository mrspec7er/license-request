package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ControllerModule(DB *gorm.DB, Memcache *redis.Client) func(chi.Router) {
	cs := Consumer{
		Service: Service{
			DB: DB,
		},
	}

	go cs.Load()

	ct := Controller{
		Service: Service{
			DB: DB,
		},
	}

	u := Middleware{
		Util: &Util{
			Memcache: Memcache,
		},
	}

	return func(r chi.Router) {
		r.With(u.Authorize()).Get("/{id}", ct.GetOne)
		r.With(u.Authorize()).Post("/", ct.Create)
		r.With(u.Authorize()).Delete("/", ct.Delete)
	}
}
