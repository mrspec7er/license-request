package src

import (
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/license-request/services/utility/dto"
)

type FormController struct {
	Service FormService
}

func (c FormController) GetAll(w http.ResponseWriter, r *http.Request) {
	forms := &dto.Form{}

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
