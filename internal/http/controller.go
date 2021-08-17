package http

import (
	"net/http"
	"net/url"
	"scheduler/pkg/router"
	"strings"
)

type ISchedulerController interface {
	List(w http.ResponseWriter, r *http.Request, p *url.Values)
	Show(w http.ResponseWriter, r *http.Request, p *url.Values)
	Create(w http.ResponseWriter, r *http.Request, p *url.Values)
	Update(w http.ResponseWriter, r *http.Request, p *url.Values)
	Delete(w http.ResponseWriter, r *http.Request, p *url.Values)
	Init(r *router.Router)
}

type Controller struct {
	Router                  *router.Router
	ScheduleEventController ISchedulerController
}

func NewController(router *router.Router, scheduleController ISchedulerController) *Controller {
	return &Controller{
		Router: router,
		ScheduleEventController: scheduleController,
	}
}


func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handleMethod, err := c.Router.GetHandleFunctionByRoute(strings.ToUpper(r.Method), r.RequestURI)
	if err != nil {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
	handleMethod(w, r, &c.Router.UrlParams)
}

func (c *Controller) Init() {
	c.ScheduleEventController.Init(c.Router)
}