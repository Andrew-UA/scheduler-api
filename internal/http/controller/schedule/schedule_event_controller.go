package schedule

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"scheduler/pkg/router"
	"strconv"
)

type Controller struct{
	ScheduleService service.IScheduleService
	UserService service.IUserService
}

func NewController(schedule service.IScheduleService, user service.IUserService) *Controller {
	return &Controller{
		ScheduleService: schedule,
		UserService: user,
	}
}

func (c Controller) Init(r *router.Router) {
	r.GET("/schedule-events", c.List)
	r.GET("/schedule-events/{id}", c.Show)
	r.POST("/schedule-events", c.Create)
	r.PUT("/schedule-events/{id}", c.Update)
	r.DELETE("/schedule-events/{id}", c.Delete)
	r.URLMiddleware("/schedule-events", []string{
		"auth", "validation",
	})
}

func (c Controller) List(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:List")

	userId, authErr := strconv.Atoi(p.Get(router.AuthorizedUserId))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(errors.New("Unauthorized").Error()))
	}
	ctx := context.WithValue(context.Background(), "authUserID", userId)
	authUser, err := c.UserService.Show(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
	params := make(map[string]string)
	interval := p.Get("interval")
	if interval != "" {
		params["interval"] = interval
	}


	scheduleEvents, sErr := c.ScheduleService.List(ctx, params)
	if sErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sErr.Error()))
		return
	}

	timezone := p.Get("timezone")
	if timezone == "" {
		timezone = authUser.Timezone
	}
	jsonEvents := make([]model.ScheduleEventJson, len(scheduleEvents))
	for key, scheduleEvent := range scheduleEvents {
		scheduleEventJson, err := scheduleEvent.ToScheduleEventJson(timezone)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(err.Error()))
			return
		}
		jsonEvents[key] = *scheduleEventJson
	}

	response, _ := json.Marshal(jsonEvents)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Show(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Show")

	userId, authErr := strconv.Atoi(p.Get(router.AuthorizedUserId))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(errors.New("Unauthorized").Error()))
	}
	ctx := context.WithValue(context.Background(), "authUserID", userId)
	authUser, err := c.UserService.Show(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	scheduleEvent, sErr := c.ScheduleService.Show(ctx, id)
	if sErr != nil {

		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sErr.Error()))
		return
	}
	timezone := p.Get("timezone")
	if timezone == "" {
		timezone = authUser.Timezone
	}
	response, _ := scheduleEvent.Marshal(timezone)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Create")

	userId, authErr := strconv.Atoi(p.Get(router.AuthorizedUserId))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(errors.New("Unauthorized").Error()))
	}
	ctx := context.WithValue(context.Background(), "authUserID", userId)
	authUser, err := c.UserService.Show(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
	mJson := model.ScheduleEventJson{}
	d :=json.NewDecoder(r.Body)
	dErr := d.Decode(&mJson)
	if dErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	timezone := p.Get("timezone")
	if timezone == "" {
		timezone = authUser.Timezone
	}
	m := model.ScheduleEvent{}
	mErr := m.FromScheduleEventJson(mJson, timezone)
	if mErr != nil {
		response, _ := json.Marshal(mErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	m, sErr := c.ScheduleService.Create(ctx, m)
	if sErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	response, _ := m.Marshal(timezone)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Update")

	userId, authErr := strconv.Atoi(p.Get(router.AuthorizedUserId))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(errors.New("Unauthorized").Error()))
	}
	ctx := context.WithValue(context.Background(), "authUserID", userId)
	authUser, err := c.UserService.Show(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	mJson := model.ScheduleEventJson{}
	d :=json.NewDecoder(r.Body)
	dErr := d.Decode(&mJson)
	if dErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	timezone := p.Get("timezone")
	if timezone == "" {
		timezone = authUser.Timezone
	}
	m := model.ScheduleEvent{}
	mErr := m.FromScheduleEventJson(mJson, timezone)
	if mErr != nil {
		response, _ := json.Marshal(mErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	m, sErr := c.ScheduleService.Update(ctx, id, m)
	if sErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	response, _ := m.Marshal(timezone)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Delete")

	userId, authErr := strconv.Atoi(p.Get(router.AuthorizedUserId))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(errors.New("Unauthorized").Error()))
	}
	ctx := context.WithValue(context.Background(), "authUserID", userId)
	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	sErr := c.ScheduleService.Delete(ctx, id)
	if sErr != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sErr.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write(nil)
	if err != nil {
		log.Fatal(err)
	}
}
