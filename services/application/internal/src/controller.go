package src

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request-utility/response"
)

type Controller struct {
	Service  Service
	Publish  Publisher
	Response response.ResponseJSON
}

func (c Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")
	app := dto.Form{}

	user := r.Context().Value(dto.UserContextKey).(dto.User)
	status, err := c.Service.ApplicationAccessGuard(number, user)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	status, err = c.Service.GetOne(&app, number)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, app, nil)
}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")
	apps := []*dto.Application{}

	status, err := c.Service.GetAll(&apps, number)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, apps, nil)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	app := &dto.Application{}
	user := r.Context().Value(dto.UserContextKey).(dto.User)

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	app.UserID = user.ID

	err := c.Publish.Create(app, user.ID)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Create form submitted")
}

func (c *Controller) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	app := &ChangeStatusInput{}
	user := r.Context().Value(dto.UserContextKey).(dto.User)

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	err := c.Publish.UpdateStatus(app, user.ID)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Update status submitted")
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	app := &dto.Application{}
	user := r.Context().Value(dto.UserContextKey).(dto.User)

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	err := c.Publish.Delete(app, user.ID)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Delete form submitted")
}
