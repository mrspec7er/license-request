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

	c := Controller{
		Service: Service{
			DB: DB,
			Util: &Util{
				Memcache: memcache,
			},
		},
	}
	u := Middleware{
		Util: &Util{
			Memcache: memcache,
		},
	}

	return func(r chi.Router) {
		r.With(u.Authorize()).Get("/{number}", c.GetOne)
		r.With(u.Authorize()).Get("/", c.GetAll)
		r.With(u.Authorize()).Post("/", c.Create)
		r.With(u.Authorize()).Put("/status", c.UpdateStatus)
		r.With(u.Authorize()).Delete("/", c.Delete)
	}
}
