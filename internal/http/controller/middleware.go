package controller

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"net/url"
	"scheduler/internal/helpers"
	"scheduler/internal/repository"
	"scheduler/internal/service"
	pkg_auth "scheduler/pkg/auth"
	"scheduler/pkg/logger"
	"scheduler/pkg/router"
	"strconv"
	"strings"
	"time"
)

type Middleware struct {
	Logger   logger.Logger
	UserRepo repository.IUser
}

func (m *Middleware) AuthCheck(next router.HandleFunction) router.HandleFunction {
	return func(w http.ResponseWriter, r *http.Request, p *url.Values) {
		m.Logger.Debugf("Hello from Auth middleware IN")

		h := r.Header.Get("Authorization")
		if h == "" {
			m.unauthorizedError(w, nil)
			return
		}

		authorizedUserIDStr, err := pkg_auth.NewTokenManager("").Parse(strings.TrimPrefix(h, "Bearer "))
		if err != nil {
			m.unauthorizedError(w, err)
			return
		}
		authorizedUserID, err := strconv.Atoi(authorizedUserIDStr)
		if err != nil {
			m.unauthorizedError(w, err)
			return
		}
		user, err := m.UserRepo.FindByID(authorizedUserID)
		if err != nil {
			m.unauthorizedError(w, err)
			return
		}

		ctx := helpers.SetUserToContext(user, r.Context())
		r = r.WithContext(ctx)
		next(w, r, p)

		m.Logger.Debugf("Hello from Auth middleware OUT")
	}
}

func (m *Middleware) Metrics(next router.HandleFunction) router.HandleFunction {
	return func(w http.ResponseWriter, r *http.Request, p *url.Values) {
		m.Logger.Debugf("Hello from metricsService middleware IN")
		begin := time.Now()
		next(w, r, p)

		end := time.Now()
		mData := service.Metric{
			Method:   r.Method,
			Route:    r.URL.Path,
			Duration: end.Sub(begin),
		}

		metricService := service.GetMetricsServiceInstance()
		metricService.Add(mData)
		m.Logger.Debugf("Method: " + mData.Method + " URL: " + mData.Route + " Duration: " + mData.Duration.String())
		m.Logger.Debugf("Hello from metricsService middleware OUT")
	}
}

func (m *Middleware) Validation(next router.HandleFunction) router.HandleFunction {
	return func(w http.ResponseWriter, r *http.Request, p *url.Values) {
		m.Logger.Debugf("Hello from Validation middleware IN")

		next(w, r, p)

		m.Logger.Debugf("Hello from Validation middleware OUT")
	}
}

func (m *Middleware) GetList() map[string]router.MiddlewareFunction {
	return map[string]router.MiddlewareFunction{

		"auth":       m.AuthCheck,
		"validation": m.Validation,
		"metrics":    m.Metrics,
	}
}

func (m *Middleware) unauthorizedError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	switch err.(type) {
	case *jwt.ValidationError:
		w.Write([]byte(err.Error()))
	default:
		w.Write([]byte("Unauthorized"))
	}
	return
}
