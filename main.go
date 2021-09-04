package main

import (
	"github.com/michaelwp/dragon/dragon"
	"github.com/michaelwp/dragon/handlers"
	"log"
)

func main() {
	r := dragon.NewRouter()

	r.POST("/api/v1/users", handlers.Create)
	r.GET("/api/v1/users", handlers.List)

	log.Fatal(r.Run(":8090"))
}