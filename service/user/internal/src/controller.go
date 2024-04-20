package auth

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/markbates/goth/gothic"
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
	user, err := c.Service.SaveUserSessions(w, r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(w, user)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	c.Service.RemoveUserSessions(w, r)
	w.Header().Set("Location", "/api/auth/index")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
