package src

import (
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/license-request-utility/dto"
)

type Controller struct {
	Service Service
	Publish Publisher
}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
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

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

	form := &dto.Form{}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	err := c.Publish.Create(*form, "123")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"message": "Form created!"})
}
