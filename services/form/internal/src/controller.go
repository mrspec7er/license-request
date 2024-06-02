package src

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 400, err)
		return
	}

	forms := &dto.Form{}

	status, err := c.Service.GetOne(forms, uint(id))
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, forms, nil)
}

func (c Controller) GetAll(w http.ResponseWriter, r *http.Request) {

	forms := []*dto.Form{}

	status, err := c.Service.GetAll(&forms)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, forms, nil)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	form := &dto.Form{}
	user := r.Context().Value(dto.UserContextKey).(dto.User)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	err := c.Publish.Create(*form, user.ID)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Create form submitted")
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	form := &dto.Form{}
	user := r.Context().Value(dto.UserContextKey).(dto.User)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	err := c.Publish.Delete(*form, user.ID)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Delete form submitted")
}
