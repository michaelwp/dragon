package main

import (
	"log"
)

func main()  {
	r := NewRouter()
	api := r.Group("/api/v1")
	api.GET("/home", Home)
	home := api.Group("/home")
	home.POST("/create", Create)

	log.Fatal(r.Run(":8090"))
}

