package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"scheduler/pkg/router"
	"strings"
)

type IMiddleware interface {
	AuthCheck(next router.HandleFunction) router.HandleFunction
	Metrics(next router.HandleFunction) router.HandleFunction
	Validation(next router.HandleFunction) router.HandleFunction
	GetList() map[string]router.MiddlewareFunction
}

type ISchedulerController interface {
	List(w http.ResponseWriter, r *http.Request, p *url.Values)
	Show(w http.ResponseWriter, r *http.Request, p *url.Values)
	Create(w http.ResponseWriter, r *http.Request, p *url.Values)
	Update(w http.ResponseWriter, r *http.Request, p *url.Values)
	Delete(w http.ResponseWriter, r *http.Request, p *url.Values)
}

type IUserController interface {
	Update(w http.ResponseWriter, r *http.Request, p *url.Values)
}

type IAuthController interface {
	Login(w http.ResponseWriter, r *http.Request, p *url.Values)
}

type IMetricController interface {
	List(w http.ResponseWriter, r *http.Request, p *url.Values)
}

type Controller struct {
	Router                  *router.Router
	Middleware              IMiddleware
	ScheduleEventController ISchedulerController
	UserController          IUserController
	AuthController          IAuthController
	MetricController        IMetricController
}

func NewController(
	router *router.Router,
	middleware IMiddleware,
	scheduleController ISchedulerController,
	userController IUserController,
	authController IAuthController,
	metricsController IMetricController,
) *Controller {
	return &Controller{
		Router:                  router,
		Middleware:              middleware,
		ScheduleEventController: scheduleController,
		UserController:          userController,
		AuthController:          authController,
		MetricController:        metricsController,
	}
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Remote addr: ", r.RemoteAddr)
	if r.Method == http.MethodOptions {
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		return
	}

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
	// Register middleware
	if c.Middleware != nil {
		c.Router.RegisterMiddleware(c.Middleware.GetList())
	}

	// Register routes
	// SCHEDULE EVENTS
	if c.ScheduleEventController != nil {
		c.Router.GET("/schedule-events", c.ScheduleEventController.List)
		c.Router.GET("/schedule-events/{id}", c.ScheduleEventController.Show)
		c.Router.POST("/schedule-events", c.ScheduleEventController.Create)
		c.Router.PUT("/schedule-events/{id}", c.ScheduleEventController.Update)
		c.Router.DELETE("/schedule-events/{id}", c.ScheduleEventController.Delete)
		c.Router.URLMiddleware("/schedule-events", []string{
			"metrics", "auth", "validation",
		})
	}

	// USERS
	if c.UserController != nil {
		c.Router.PUT("/users/{id}", c.UserController.Update)
		c.Router.URLMiddleware("/users", []string{
			"metrics", "auth",
		})
	}

	// AUTH
	if c.AuthController != nil {
		c.Router.POST("/login", c.AuthController.Login)
		c.Router.URLMiddleware("/login", []string{
			"metrics",
		})
	}

	// METRICS
	if c.MetricController != nil {
		c.Router.GET("/metrics", c.MetricController.List)
	}
}
