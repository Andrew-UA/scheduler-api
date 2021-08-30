package user

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

type Controller struct {
	UserService service.IUserService
}

func NewController(user service.IUserService) *Controller {
	return &Controller{
		UserService: user,
	}
}

func (c Controller) Init(r *router.Router) {
	r.PUT("/users/{id}", c.Update)
	r.URLMiddleware("/users", []string{
		"auth",
	})
}

func (c Controller) Update(w http.ResponseWriter, r *http.Request, p *url.Values)  {
	fmt.Println("UserController:Show")

	userId, authErr := strconv.Atoi(p.Get(router.AuthorizedUserId))
	if authErr != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(errors.New("Unauthorized").Error()))
		return
	}
	ctx := context.WithValue(context.Background(), "authUserID", userId)
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
	user, sErr := c.UserService.Update(ctx, id, u)
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
		log.Fatal(err)
	}
}

