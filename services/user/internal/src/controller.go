package src

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/mrspec7er/license-request/services/utility/dto"
)

type AuthController struct {
	Service AuthService
}

func (c *AuthController) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, false)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (c *AuthController) Callback(w http.ResponseWriter, r *http.Request) {
	user := &goth.User{}
	err := c.Service.StoreUserSessions(w, r, user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(w, user)
}

func (c *AuthController) Info(w http.ResponseWriter, r *http.Request) {
	authKey := chi.URLParam(r, "authKey")
	user := &dto.User{}

	err := c.Service.RetrieveUserSessions(w, r, authKey, user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	c.Service.RemoveUserSessions(w, r)
	w.Header().Set("Location", "/api/auth/index")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
