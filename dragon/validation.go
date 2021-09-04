package dragon

import (
	"errors"
	"net/http"
)

func (r *router) isMethodAllowed(rw http.ResponseWriter, req *http.Request) error {
	var errorMessage = "method not allowed"

	if !isMatching(r.Method, req.Method) {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte(errorMessage))
		return errors.New(errorMessage)
	}

	return nil
}
