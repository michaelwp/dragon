package dragon

import (
	"net/http"
)

var routers = make([]*router, 0)

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

func updateRouter(method string, hh HandlerFunc, rr *router) {
	var h handler

	h.Handler = hh
	h.Methods = method

	rr.Handler = append(rr.Handler, h)
}

func setupRouter(method string, hh HandlerFunc, path string) {
	isTrue, rr := isRouterRegistered(path)
	if isTrue {
		// update specific router
		updateRouter(method, hh, rr)
		return
	}

	// register new router
	registerRouter(method, hh, path)

	// setup http handle with routers as the handler
	http.Handle(path, &router{})
}

func registerRouter(method string, hh HandlerFunc, path string) {
	var r router
	var h handler

	// form handler that will register to the router
	h.Handler = hh
	h.Methods = method

	// register params if any
	params := registerParams(path)

	// register new router
	r.Path = path
	r.Handler = append(r.Handler, h)
	r.Params = params

	routers = append(routers, &r)
}

/* METHODS */

func (r *router) GET(path string, hh HandlerFunc) {
	setupRouter(http.MethodGet, hh, r.RouterGroup+path)
}

func (r *router) POST(path string, hh HandlerFunc) {
	setupRouter(http.MethodPost, hh, r.RouterGroup+path)
}
