package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"scheduler/internal/model"
	"scheduler/internal/service"
	"scheduler/pkg/logger"
	"strconv"
)

type UserController struct {
	UserService service.IUserService
	Logger      logger.Logger
}

func NewUserController(user service.IUserService, logger logger.Logger) *UserController {
	return &UserController{
		UserService: user,
		Logger:      logger,
	}
}

func (c UserController) Update(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("UserController:Show")

	id, convErr := strconv.Atoi(p.Get("id"))
	if convErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(convErr.Error()))
		return
	}

	u := model.User{}
	d := json.NewDecoder(r.Body)
	dErr := d.Decode(&u)
	if dErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(dErr.Error()))
		return
	}
	user, sErr := c.UserService.Update(r.Context(), id, u)
	if sErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(sErr.Error()))
		return
	}

	response, mErr := json.Marshal(user)
	if mErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(mErr.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(response)
	if err != nil {
		c.Logger.Error(err)
	}
}
