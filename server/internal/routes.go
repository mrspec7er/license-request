package internal

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrspec7er/license-request/server/internal/module/application"
	"github.com/mrspec7er/license-request/server/internal/module/auth"
	"github.com/mrspec7er/license-request/server/internal/module/form"
)

func (s Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello There!"})
	})

	router.Route("/forms", form.Module(s.DB))
	router.Route("/apps", application.Module(s.DB))
	router.Route("/auth", auth.Module(s.DB))

	return router
}
