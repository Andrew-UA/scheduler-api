package schedule

import (
	"encoding/json"
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
	Service service.IScheduleService
}

func NewController(s service.IScheduleService) *Controller {
	return &Controller{
		Service: s,
	}
}

func (c Controller) Init(r *router.Router) {
	r.GET("/schedule-events", c.List)
	r.GET("/schedule-events/{id}", c.Show)
	r.POST("/schedule-events", c.Create)
	r.PUT("/schedule-events/{id}", c.Update)
	r.DELETE("/schedule-events/{id}", c.Delete)
}

func (c Controller) List(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:List")
	params := make(map[string]string)
	interval := p.Get("interval")
	if interval != "" {
		params["interval"] = interval
	}

	scheduleEvents, sErr := c.Service.List(params)
	if sErr != nil {
		response, _ := json.Marshal(sErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	jsonEvents := make([]model.ScheduleEventJson, len(scheduleEvents))
	for key, scheduleEvent := range scheduleEvents {
		jsonEvents[key] = *scheduleEvent.ToScheduleEventJson()
	}

	response, _ := json.Marshal(jsonEvents)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Show(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Show")

	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	scheduleEvent, sErr := c.Service.Show(id)
	if sErr != nil {

		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(sErr.Error()))
		return
	}
	response, _ := scheduleEvent.Marshal()
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Create")

	mJson := model.ScheduleEventJson{}
	d :=json.NewDecoder(r.Body)
	dErr := d.Decode(&mJson)
	if dErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	m := model.ScheduleEvent{}
	mErr := m.FromScheduleEventJson(mJson)
	if mErr != nil {
		response, _ := json.Marshal(mErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	m, sErr := c.Service.Create(m)
	if sErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	response, _ := m.Marshal()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Update")

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

	m := model.ScheduleEvent{}
	mErr := m.FromScheduleEventJson(mJson)
	if mErr != nil {
		response, _ := json.Marshal(mErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	m, sErr := c.Service.Update(id, m)
	if sErr != nil {
		response, _ := json.Marshal(dErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	response, _ := m.Marshal()
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) Delete(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("ScheduleEventController:Delete")
	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		response, _ := json.Marshal(convErr)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
		return
	}

	sErr := c.Service.Delete(id)
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
