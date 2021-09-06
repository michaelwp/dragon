package main

import (
	"errors"
	"github.com/michaelwp/dragon"
	"net/http"
)

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    interface{}
}

func Authorize() dragon.HandlerFunc {
	var errMessage errorMessage

	errMessage.Code = 0
	errMessage.Message = "Unauthorized"

	return func(d *dragon.Dragon) error {
		if d.RequestHeader.Get("token") != "token123456" {
			d.ResponseJSON(http.StatusUnauthorized, errMessage)
			return errors.New(errMessage.Message)
		}

		d.SetContext("token", d.RequestHeader.Get("token"))
		d.SetContext("userID", 1)

		return d.Next()
	}
}
