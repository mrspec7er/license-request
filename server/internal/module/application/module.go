package application

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Module(db *gorm.DB) func(chi.Router) {
	c := ApplicationController{
		Service: ApplicationService{
			DB: db,
		},
	}

	return func(r chi.Router) {
		r.Get("/{number}", c.GetAll)
	}
}
