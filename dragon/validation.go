package dragon

import (
	"errors"
	"net/http"
)

const (
	PathNotFound     = "path not found"
	MethodNotAllowed = "method not allowed"
)

func isPathExist(rw http.ResponseWriter, req *http.Request) (*handler, error) {
	for _, rr := range routers {
		if isMatching(req.URL.Path, rr.Path) {
			hh, err := rr.isMethodAllowed(rw, req)
			return hh, err
		}
	}

	// return error not found as default response
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(PathNotFound))
	return nil, errors.New(PathNotFound)
}

func (r *router) isMethodAllowed(rw http.ResponseWriter, req *http.Request) (*handler, error) {
	for _, hh := range r.Handler {
		if isMatching(req.Method, hh.Methods) {
			return &hh, nil
		}
	}

	// return error method not allowed as default response
	rw.WriteHeader(http.StatusMethodNotAllowed)
	rw.Write([]byte(MethodNotAllowed))
	return nil, errors.New(MethodNotAllowed)
}
