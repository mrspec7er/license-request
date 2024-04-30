package src

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func ControllerModule(DB *gorm.DB) func(chi.Router) {
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

	return func(r chi.Router) {
		r.Get("/", ct.GetAll)
		r.Post("/", ct.Create)
	}
}
