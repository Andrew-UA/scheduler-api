package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"scheduler/internal/service"
	"scheduler/pkg/logger"
)

type InputBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthController struct {
	Service service.IAuthService
	Logger  logger.Logger
}

func NewAuthController(s service.IAuthService, logger logger.Logger) *AuthController {
	return &AuthController{
		Service: s,
		Logger:  logger,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request, p *url.Values) {
	c.Logger.Debugf("AuthController:Login")
	input := &InputBody{}
	d := json.NewDecoder(r.Body)
	dErr := d.Decode(input)

	if dErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(dErr.Error()))
		return
	}
	token, err := c.Service.SignIn(input.Login, input.Password)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
