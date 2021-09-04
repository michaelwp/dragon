package dragon

import (
	"log"
	"net/http"
)

func NewRouter() Router {
	return &router{}
}

func logging(rw http.ResponseWriter, req *http.Request) {
	log.Printf("%s - %s %s", req.RemoteAddr, req.Method, req.URL.Path)
}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// validating http method requested
	hh, err := isPathExist(rw, req)
	if err != nil {
		return
	}

	// checking middleware
	err = middlewareChecking(rw, req)
	if err != nil {
		return
	}

	// simple request logging
	logging(rw, req)

	// update handler http response writer & request
	hh.Handler(rw, req)
}

func (r *router) Run(address string) error {
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}
