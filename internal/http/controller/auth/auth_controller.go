package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"scheduler/internal/service"
	"scheduler/pkg/router"
)

type InputBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Controller struct {
	Service service.IAuthService
}

func NewController(s service.IAuthService) *Controller {
	return &Controller{
		Service: s,
	}
}

func (c Controller) Init(r *router.Router) {
	r.POST("/login", c.Login)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request, p *url.Values) {
	fmt.Println("AuthController:Login")
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
	if  err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
