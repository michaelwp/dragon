package dragon

import (
	"net/http"
)

var routers = make([]*router, 0)
var handlers = make([]*handler, 0)

func (r router) Group(path string) router {
	r.RouterGroup = r.RouterGroup + path
	return r
}

func isRouterRegistered(path string) (bool, *router) {
	for _, rr := range routers {
		if isMatching(path, rr.Path) {
			return true, rr
		}
	}
	return false, nil
}

func registerRouter(method *string, hh http.HandlerFunc, path string) {
	// register new handler that will register to new router
	h := &handler{
		Handler: hh,
		Methods: method,
	}
	handlers = append(handlers, h)

	// register new router
	r := &router{
		Path:    path,
		Handler: handlers,
	}
	routers = append(routers, r)

	// setup http handle with routers as the handler
	http.Handle(path, r)
}

func setupRouter(method string, hh http.HandlerFunc, path string) {
	isTrue, rr := isRouterRegistered(path)
	if isTrue {
		updateRouter(&method, hh, rr)
		return
	}
	registerRouter(&method, hh, path)
}

func updateRouter(method *string, hh http.HandlerFunc, rr *router) {
	h := &handler{
		Handler: hh,
		Methods: method,
	}
	rr.Handler = append(rr.Handler, h)
}

/* METHODS */

func (r *router) GET(path string, hh http.HandlerFunc) {
	setupRouter(http.MethodGet, hh, r.RouterGroup+path)
}

func (r *router) POST(path string, hh http.HandlerFunc) {
	setupRouter(http.MethodPost, hh, r.RouterGroup+path)
}
