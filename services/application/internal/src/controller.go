package src

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/license-request-utility/dto"
)

type Controller struct {
	Service Service
}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")
	app := dto.Form{}

	user := r.Context().Value(dto.UserContextKey).(dto.User)

	status, err := c.Service.ApplicationAccessGuard(number, user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	status, err = c.Service.GetOne(&app, number)
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
