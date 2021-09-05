package main

import (
	"errors"
	"github.com/michaelwp/dragon"
	"net/http"
)

func Authorize() func(*dragon.Dragon) error {
	var errorMessage = struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    interface{}
	}{
		Code:    0,
		Message: "Unauthorized",
	}
	return func(d *dragon.Dragon) error {
		if d.RequestHeader.Get("token") != "token" {
			err := d.ResponseJSON(http.StatusUnauthorized, errorMessage)
			if err != nil {
				return err
			}
			return errors.New(errorMessage.Message)
		}
		return nil
	}
}
