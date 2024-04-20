package auth

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Module(db *gorm.DB) func(chi.Router) {
	store := AuthInit()

	c := AuthController{
		Service: AuthService{
			DB:    db,
			Store: store,
		},
	}

	return func(r chi.Router) {
		r.Get("/", c.Index)
		r.Get("/login", c.Login)
		r.Get("/callback", c.Callback)
		r.Get("/logout", c.Logout)
	}
}