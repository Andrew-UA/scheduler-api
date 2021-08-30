package middleware

import (
	"scheduler/pkg/router"
)

func GetList() map[string]router.MiddlewareFunction {
	//TODO: change
	return map[string]router.MiddlewareFunction{
		"auth":       AuthCheck,
		"validation": Validation,
	}
}
