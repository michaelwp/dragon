package dragon

import (
	"log"
	"net/http"
)

func NewRouter() Router {
	return &router{}
}

func logging(req *http.Request) {
	log.Printf("%s - %s %s", req.RemoteAddr, req.Method, req.URL.Path)
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

	// simple request logging
	logging(req)

	// mapping params
	mParams := rr.mapParamsUrl()

	// setup dragon handler
	d := &Dragon{
		ResponseWriter: rw,
		Request:        req,
		Context:        req.Context(),
		Params:         mParams,
		Headers:        req.Header,
		Path:           req.URL.Path,
		Body:           req.Body,
		Query:          req.URL.Query(),
		RemoteAddress:  req.RemoteAddr,
	}

	// update handler http response writer & request
	hh.Handler(d)
}

func (r *router) Run(address string) error {
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, r)
	if err != nil {
		return err
	}

	return nil
}
