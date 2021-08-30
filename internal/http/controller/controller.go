package controller

import (
	"net/http"
	"net/url"
	"scheduler/internal/http/middleware"
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

type IUserController interface {
	Update(w http.ResponseWriter, r *http.Request, p *url.Values)
	Init(r *router.Router)
}

type IAuthController interface {
	Login(w http.ResponseWriter, r *http.Request, p *url.Values)
	Init(r *router.Router)
}

type Controller struct {
	Router                  *router.Router
	ScheduleEventController ISchedulerController
	UserController          IUserController
	AuthController          IAuthController
}

func NewController(
	router *router.Router,
	scheduleController ISchedulerController,
	userController IUserController,
	authController IAuthController,
) *Controller {
	return &Controller{
		Router:                  router,
		ScheduleEventController: scheduleController,
		UserController:          userController,
		AuthController:          authController,
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
	c.Router.RegisterMiddleware(middleware.GetList())
	c.ScheduleEventController.Init(c.Router)
	c.AuthController.Init(c.Router)
	c.UserController.Init(c.Router)
}
