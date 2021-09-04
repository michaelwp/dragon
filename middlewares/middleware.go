package middlewares

import (
	"errors"
	"net/http"
)

func Logging() func(rw http.ResponseWriter, req *http.Request) error {
	var errorMessage = "authentication error"
	return func(rw http.ResponseWriter, req *http.Request) error {
		if req.Header.Get("token") != "token" {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			rw.Write([]byte(errorMessage))
			return errors.New(errorMessage)
		}
		return nil
	}
}
