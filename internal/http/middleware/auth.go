package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
	"net/url"
	pkg_auth "scheduler/pkg/auth"
	"scheduler/pkg/router"
	"strings"
)

func AuthCheck(next router.HandleFunction) router.HandleFunction {
	return func(w http.ResponseWriter, r *http.Request, p *url.Values) {
		fmt.Println("Hello from Auth middleware IN")

		h := r.Header.Get("Authorization")
		if h == "" {
			notAuthorizedError(w, nil)
			return
		}

		authorizedUserID, err := pkg_auth.NewTokenManager("").Parse(strings.TrimPrefix(h, "Bearer "))
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError:
				notAuthorizedError(w, err)
			default:
				notAuthorizedError(w, nil)
			}

			return
		}
		p.Add(router.AuthorizedUserId, authorizedUserID)

		next(w, r, p)

		fmt.Println("Hello from Auth middleware OUT")
	}
}

func notAuthorizedError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Unauthorized"))
	}
	return
}
