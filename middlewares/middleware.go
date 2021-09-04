package middlewares

import (
	"errors"
	"log"
	"net/http"
)

func Authorize() func(rw http.ResponseWriter, req *http.Request) error {
	var errorMessage = "authentication error"
	return func(rw http.ResponseWriter, req *http.Request) error {
		if req.Header.Get("token") != "token" {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte(errorMessage))
			return errors.New(errorMessage)
		}
		return nil
	}
}

func Logging() func(rw http.ResponseWriter, req *http.Request) error {
	return func(rw http.ResponseWriter, req *http.Request) error {
		log.Println(req.Method, req.URL.Path)
		return nil
	}
}