package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/license-request/services/form/internal/db"
)

func ControllerModule(DB *db.Conn, Memcache *db.CacheClient) func(chi.Router) {
	cs := Consumer{
		Service: Service{
			Store: db.FormRepository{
				DB: DB,
			},
		},
	}

	go cs.Load()

	ct := Controller{
		Service: Service{
			Store: db.FormRepository{
				DB: DB,
			},
		},
	}

	u := Middleware{
		Cache: db.RedisRepository{
			Cache: Memcache,
		},
	}

	return func(r chi.Router) {
		r.With(u.Authorize()).Get("/", ct.GetAll)
		r.With(u.Authorize()).Get("/{id}", ct.GetOne)
		r.With(u.Authorize()).Post("/", ct.Create)
		r.With(u.Authorize()).Delete("/", ct.Delete)
	}
}
