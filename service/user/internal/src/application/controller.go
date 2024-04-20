package application

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/license-request/service/user/internal/db"
)

type ApplicationController struct {
	Service ApplicationService
}

func (c ApplicationController) GetAll(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")
	app := db.Form{}

	status, err := c.Service.GetOne(&app, number)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(app)
}
