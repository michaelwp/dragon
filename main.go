package main

import (
	"github.com/michaelwp/dragon/dragon"
	"github.com/michaelwp/dragon/handlers"
	"log"
)

func main() {
	r := dragon.NewRouter()
	r.GET("/api/v1/home", handlers.Home)

	r.GET("/api/v1/users", handlers.List)
	r.POST("/api/v1/users", handlers.Create)

	log.Fatal(r.Run(":8090"))
}