package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/license-request/services/application/internal/db"
)

func Module(DB *db.Conn, memcache *db.CacheClient) func(chi.Router) {
	cs := Consumer{
		Service: Service{
			Store: db.AppRepository{
				DB: DB,
			},
		},
	}

	go cs.Load()

	c := Controller{
		Service: Service{
			Store: db.AppRepository{
				DB: DB,
			},
		},
	}
	u := Middleware{
		Cache: db.RedisRepository{Cache: memcache},
	}

	return func(r chi.Router) {
		r.With(u.Authorize()).Get("/{number}", c.GetOne)
		r.With(u.Authorize()).Get("/", c.GetAll)
		r.With(u.Authorize()).Post("/", c.Create)
		r.With(u.Authorize()).Put("/status", c.UpdateStatus)
		r.With(u.Authorize()).Delete("/", c.Delete)
	}
}
