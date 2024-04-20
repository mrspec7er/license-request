package src

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Module(db *gorm.DB) func(chi.Router) {
	c := FormController{
		Service: FormService{
			DB: db,
		},
	}

	return func(r chi.Router) {
		r.Get("/", c.GetAll)
	}
}
