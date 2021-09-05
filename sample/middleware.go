package main

import (
	"errors"
	"github.com/michaelwp/dragon"
	"net/http"
)

func Authorize() func(*dragon.Dragon) error {
	var errorMessage = "authentication error"
	return func(d *dragon.Dragon) error {
		if d.Header.Get("token") != "token" {
			d.ResponseStatus(http.StatusUnauthorized)
			d.ResponseWriter.Write([]byte(errorMessage))
			return errors.New(errorMessage)
		}
		return nil
	}
}
