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
		RequestHeader:  req.Header,
		ResponseHeader: rw.Header(),
		Path:           req.URL.Path,
		Query:          req.URL.Query(),
		RemoteAddress:  req.RemoteAddr,
	}

}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// validating http request and return updated router
	hh, rr, err := isPathExist(rw, req)
	if err != nil {
		return
	}

	// setup dragon parameter
	d := rr.setupDragonParameter(rw, req)

	// checking middleware
	err = middlewareChecking(d)
	if err != nil {
		return
	}

	// update handler http response writer & request
	err = hh.Handler(d)
	if err != nil {
		log.Fatal(err)
	}

	// simple request logging
	logging(req)
}

// Run http server, example : log.Fatal(r.Run(":8090"))
func (r *router) Run(address string) error {
	var fileSystem http.Handler

	// setup static html file if exist
	if r.FileSystem != nil {
		fileSystem = http.FileServer(http.Dir(*r.FileSystem))
	}

	// start http server
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, fileSystem)
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

// ServeHTTPFile serving http file with specific pattern.
// 	example :
// 		r := dragon.NewRouter()
//		r.ServeHTTPFile("./static")
func (r *router) ServeHTTPFile(systemFile string) {
	r.FileSystem = &systemFile
}

/* HTTP METHODS */

func (r *router) GET(path string, hh HandlerFunc) {
	setupRouter(http.MethodGet, hh, r.RouterGroup+path)
}

func (r *router) POST(path string, hh HandlerFunc) {
	setupRouter(http.MethodPost, hh, r.RouterGroup+path)
}

func (r *router) PUT(path string, hh HandlerFunc) {
	setupRouter(http.MethodPut, hh, r.RouterGroup+path)
}

func (r *router) PATCH(path string, hh HandlerFunc) {
	setupRouter(http.MethodPatch, hh, r.RouterGroup+path)
}

func (r *router) DELETE(path string, hh HandlerFunc) {
	setupRouter(http.MethodDelete, hh, r.RouterGroup+path)
}
