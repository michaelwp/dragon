package dragon

import (
	"net/http"
)



func (r *router) Group(path string) *router {
	r.RouterGroup = r.RouterGroup + path
	return r
}

func setupRoute(method string, handler http.HandlerFunc, path string) {
	r := router{
		Path:    path,
		Handler: handler,
		Method:  method,
	}

	http.Handle(path, &r)
}

/* METHODS */

func (r *router) GET(path string, handlerFunc http.HandlerFunc) {
	setupRoute(http.MethodGet, handlerFunc, r.RouterGroup+path)
}

func (r *router) POST(path string, handlerFunc http.HandlerFunc) {
	setupRoute(http.MethodPost, handlerFunc, r.RouterGroup+path)
}
