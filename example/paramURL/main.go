package main

import (
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
)

func main() {
	r := dragon.NewRouter()

	// api
	api := r.Group("/api/v1")
	
	// product
	product := api.Group("/products")
	product.GET("/:id/:type", func(d *dragon.Dragon) error {
		var resp response

		resp.Code = 1
		resp.Message = "Get Specific Product"
		resp.Data = d.Params

		return d.ResponseJSON(http.StatusOK, resp)
	})

	log.Fatal(r.Run(":8090"))
}
