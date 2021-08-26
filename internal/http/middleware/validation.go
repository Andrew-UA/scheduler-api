package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"scheduler/pkg/router"
)

func Validation(next router.HandleFunction) router.HandleFunction {
	return func(w http.ResponseWriter, r *http.Request, p *url.Values) {
		fmt.Println("Hello from Validation middleware IN")

		next(w,r,p)

		fmt.Println("Hello from Validation middleware OUT")
	}
}