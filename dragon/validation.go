package dragon

import (
	"errors"
	"net/http"
)

const (
	PageNotFound     = "page not found"
	MethodNotAllowed = "method not allowed"
)

func isPathExist(rw http.ResponseWriter, req *http.Request) (*handler, *router, error) {
	for _, rr := range routers {
		var reqPath = req.URL.Path
		var rrPath = rr.Path

		// check params
		if rr.isParamsExist() {
			rrPath = rr.formParamsUrl(reqPath)
		}

		// is route matching
		if isMatching(reqPath, rrPath) {
			hh, err := rr.isMethodAllowed(rw, req)
			return hh, rr, err
		}
	}

	// return error not found as default response
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(PageNotFound))
	return nil, nil, errors.New(PageNotFound)
}

func (r *router) isParamsExist() bool {
	if len(r.Params) > 0 {
		return true
	}

	return false
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
