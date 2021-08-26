package user

import (
	"net/http"
	"net/url"
	"scheduler/pkg/router"
)

type Controller struct {

}

func NewController() *Controller {
	return &Controller{}
}

func (c Controller) Init(r *router.Router) {
	r.PUT("/users/{id}", c.Update)
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request, p *url.Values)  {

}

