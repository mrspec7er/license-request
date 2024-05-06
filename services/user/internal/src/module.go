package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Module(DB *gorm.DB, Memcache *redis.Client) func(chi.Router) {
	cs := Consumer{
		Service: Service{
			DB: DB,
		},
	}

	go cs.Load()

	store := AuthInit()

	c := Controller{
		Service: Service{
			DB:    DB,
			Store: store,
			Util: &Utility{
				Memcache: Memcache,
			},
		},
	}

	return func(r chi.Router) {
		r.Get("/", c.Index)
		r.Get("/login", c.Login)
		r.Get("/callback", c.Callback)
		r.Get("/logout", c.Logout)
		r.Get("/info/{authKey}", c.Info)
		r.Get("/{uid}", c.GetOne)
	}
}
