package src

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request-utility/response"
)

type Controller struct {
	Service  Service
	Publish  Publisher
	Response response.ResponseJSON
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, false)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (c *Controller) Callback(w http.ResponseWriter, r *http.Request) {
	user := &goth.User{}
	err := c.Service.StoreUserSessions(w, r, user)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 400, err)
		return
	}

	userEntry := &dto.User{
		ID:            user.UserID,
		Picture:       user.AvatarURL,
		Email:         user.Email,
		VerifiedEmail: true,
		Role:          "USER",
	}

	err = c.Publish.Create(userEntry, user.UserID)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 400, err)
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(w, user)
}

func (c *Controller) Info(w http.ResponseWriter, r *http.Request) {
	authKey := chi.URLParam(r, "authKey")
	user := &dto.User{}

	err := c.Service.RetrieveUserSessions(w, r, authKey, user)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 400, err)
		return
	}
	c.Response.QuerySuccessResponse(w, nil, user, nil)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	c.Service.RemoveUserSessions(w, r)
	w.Header().Set("Location", "/api/auth/index")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (c *Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "uid")
	user := &dto.User{}

	status, err := c.Service.GetOne(user, uid)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}
	c.Response.QuerySuccessResponse(w, nil, user, nil)
}
