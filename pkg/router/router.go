package router

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type RouteError struct {
	Err string
}

func (e RouteError) Error() string {
	return e.Err
}

type HandleFunction func(w http.ResponseWriter, r *http.Request, p *url.Values)

func urlToRegexp(url string) string {
	var re = regexp.MustCompile(`{id}`)
	return re.ReplaceAllString(url, `\d+`)
}

type Router struct {
	get       map[string]HandleFunction
	post      map[string]HandleFunction
	put       map[string]HandleFunction
	delete    map[string]HandleFunction
	UrlParams url.Values
}

func NewRouter() *Router {
	return &Router{
		make(map[string]HandleFunction),
		make(map[string]HandleFunction),
		make(map[string]HandleFunction),
		make(map[string]HandleFunction),
		make(map[string][]string),
	}
}

func (r *Router) GET(url string, handler HandleFunction) {
	r.get[url] = handler
}

func (r *Router) POST(url string, handler HandleFunction) {
	r.post[url] = handler
}

func (r *Router) PUT(url string, handler HandleFunction) {
	r.put[url] = handler
}

func (r *Router) DELETE(url string, handler HandleFunction) {
	r.delete[url] = handler
}

func (r *Router) GetHandleFunctionByRoute(method string, urlString string) (HandleFunction, error) {
	var handler HandleFunction
	var urlParams map[string][]string
	var err error

	switch method {
	case http.MethodGet:
		handler, urlParams, err = findHandlerAndParseURL(r.get, urlString)
		if err != nil {
			return nil, err
		}

	case http.MethodPost:
		handler, urlParams, err = findHandlerAndParseURL(r.post, urlString)
		if err != nil {
			return nil, err
		}

	case http.MethodPut:
		handler, urlParams, err = findHandlerAndParseURL(r.put, urlString)
		if err != nil {
			return nil, err
		}

	case http.MethodDelete:
		handler, urlParams, err = findHandlerAndParseURL(r.delete, urlString)
		if err != nil {
			return nil, err
		}

	default:
		err = RouteError{
			"Bad Method",
		}
		return nil, err
	}
	r.UrlParams = urlParams

	return handler, nil
}

func findHandlerAndParseURL(routes map[string]HandleFunction, urlString string) (HandleFunction, url.Values, error) {
	u, _ := url.Parse(urlString)
	for handlerUrl, handler := range routes {

		// Clear urlString from last "/" if exist
		path := strings.TrimSuffix(u.Path, "/")

		splitUrl := strings.Split(path, "/")
		splitHandlerUrl := strings.Split(handlerUrl, "/")

		if len(splitUrl) != len(splitHandlerUrl) {
			continue
		}
		// Create reg exp from route urlString ("{id}" -> "\d+")
		reg := urlToRegexp(handlerUrl)
		isMatch, err := regexp.MatchString(reg, path)
		if err != nil {
			return nil, nil, err
		}
		if isMatch {
			// Parse urlString and query params
			urlValues, _ := url.ParseQuery(u.RawQuery)
			for i:=0; i < len(splitUrl); i++ {
				if splitUrl[i] != splitHandlerUrl[i] {
					paramName := splitHandlerUrl[i]
					paramName = strings.TrimPrefix(paramName, "{")
					paramName = strings.TrimSuffix(paramName, "}")
					paramValue := splitUrl[i]

					urlValues.Set(paramName, paramValue)
				}
			}

			return handler, urlValues, nil
		}
	}

	var err RouteError
	err = RouteError{
		"Unused URL",
	}

	return nil, nil, err
}