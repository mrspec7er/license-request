package src

import (
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/license-request/service/form/internal/db"
)

type FormController struct {
	Service FormService
}

func (c FormController) GetAll(w http.ResponseWriter, r *http.Request) {
	forms := &db.Form{}

	status, err := c.Service.GetOne(forms)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(forms)
}
