package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Module(db *gorm.DB, memcache *redis.Client) func(chi.Router) {
	store := AuthInit()

	c := Controller{
		Service: Service{
			DB:    db,
			Store: store,
			Util: &Utility{
				Memcache: memcache,
			},
		},
	}

	return func(r chi.Router) {
		r.Get("/", c.Index)
		r.Get("/login", c.Login)
		r.Get("/callback", c.Callback)
		r.Get("/logout", c.Logout)
		r.Get("/info/{authKey}", c.Info)
	}
}
