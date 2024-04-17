package form

import (
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/license-request/server/internal/dto"
)

type FormController struct {
	Service FormService
}

func (c FormController) GetAll(w http.ResponseWriter, r *http.Request) {
	forms := []*dto.Form{}

	status, err := c.Service.GetAll(&forms)
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
