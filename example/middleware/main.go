package main

import (
	"errors"
	"github.com/michaelwp/dragon"
	"log"
	"net/http"
)

type (
	response struct {
		Code    int         `json:"code" xml:"code"`
		Message string      `json:"message" xml:"message"`
		Data    interface{} `json:"data" xml:"data"`
	}

	errorMessage struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    interface{}
	}
)

func main() {
	var resp response

	r := dragon.NewRouter()

	// api
	api := r.Group("/api/v1")

	// home
	home := api.Group("/home")
	home.Use(Authorize())
	home.GET("", func(d *dragon.Dragon) error {
		resp.Message = "Hello World"
		return d.ResponseJSON(http.StatusOK, resp)
	})
	home.POST("", Authorize())

	log.Fatal(r.Run(":8090"))
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
