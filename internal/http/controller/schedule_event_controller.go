package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"scheduler/internal/helpers"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"scheduler/pkg/logger"
	"strconv"
)

type ScheduleController struct {
	ScheduleService service.IScheduleService
	UserService     service.IUserService
	Logger          logger.Logger
}

func NewScheduleController(schedule service.IScheduleService, user service.IUserService, logger logger.Logger) *ScheduleController {
	return &ScheduleController{
		ScheduleService: schedule,
		UserService:     user,
		Logger:          logger,
	}
}

func (c ScheduleController) List(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("ScheduleEventController:List")

	authUser, ctxErr := helpers.GetUserFormContext(r.Context())
	if ctxErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ctxErr.Error()))
		return
	}
	params := make(map[string]string)
	interval := p.Get("interval")
	if interval != "" {
		params["interval"] = interval
	}

	scheduleEvents, sErr := c.ScheduleService.List(r.Context(), params)
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
	_, err := w.Write(response)
	if err != nil {
		c.Logger.Error(err)
	}
}

func (c ScheduleController) Show(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("ScheduleEventController:Show")

	authUser, ctxErr := helpers.GetUserFormContext(r.Context())
	if ctxErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ctxErr.Error()))
		return
	}
	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	scheduleEvent, sErr := c.ScheduleService.Show(r.Context(), id)
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
	_, err := w.Write(response)
	if err != nil {
		c.Logger.Error(err)
	}
}

func (c ScheduleController) Create(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("ScheduleEventController:Create")

	authUser, ctxErr := helpers.GetUserFormContext(r.Context())
	if ctxErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ctxErr.Error()))
		return
	}
	mJson := model.ScheduleEventJson{}
	d := json.NewDecoder(r.Body)
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

	m, sErr := c.ScheduleService.Create(r.Context(), m)
	if sErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	response, _ := m.Marshal(timezone)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, err := w.Write(response)
	if err != nil {
		c.Logger.Error(err)
	}
}

func (c ScheduleController) Update(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("ScheduleEventController:Update")

	authUser, ctxErr := helpers.GetUserFormContext(r.Context())
	if ctxErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ctxErr.Error()))
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
	d := json.NewDecoder(r.Body)
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

	m, sErr := c.ScheduleService.Update(r.Context(), id, m)
	if sErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	response, _ := m.Marshal(timezone)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(response)
	if err != nil {
		c.Logger.Error(err)
	}
}

func (c ScheduleController) Delete(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("ScheduleEventController:Delete")

	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	sErr := c.ScheduleService.Delete(r.Context(), id)
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
		c.Logger.Error(err)
	}
}
