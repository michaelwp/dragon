package dragon

import (
	"log"
	"net/http"
)

var routers = make([]*router, 0)

func NewRouter() Router {
	return &router{}
}

func logging(req *http.Request) {
	log.Printf("%s - %s %s", req.RemoteAddr, req.Method, req.URL.Path)
}

func (r *router) setupDragonParameter(rw http.ResponseWriter, req *http.Request) *Dragon {
	return &Dragon{
		ResponseWriter: rw,
		Request:        req,
		Context:        req.Context(),
		Params:         r.mapParamsUrl(),
		Headers:        req.Header,
		Path:           req.URL.Path,
		Query:          req.URL.Query(),
		RemoteAddress:  req.RemoteAddr,
	}

}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// validating http method requested
	hh, rr, err := isPathExist(rw, req)
	if err != nil {
		return
	}

	// checking middleware
	err = middlewareChecking(rw, req)
	if err != nil {
		return
	}

	// setup dragon parameter
	dragon := rr.setupDragonParameter(rw, req)

	// update handler http response writer & request
	hh.Handler(dragon)

	// simple request logging
	logging(req)
}

func (r *router) Run(address string) error {
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, r)
	if err != nil {
		return err
	}

	return nil
}

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
	params := registerParamsUrl(path)

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
