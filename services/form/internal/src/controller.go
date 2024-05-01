package src

import (
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/license-request-utility/dto"
	"github.com/mrspec7er/license-request-utility/response"
)

type Controller struct {
	Service  Service
	Publish  Publisher
	Response response.ResponseJSON
}

func (c Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	forms := &dto.Form{}

	status, err := c.Service.GetOne(forms)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, forms, nil)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	form := &dto.Form{}

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	err := c.Publish.Create(*form, "123")
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Create form submitted")
}
