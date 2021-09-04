package main

import (
	"github.com/michaelwp/dragon/dragon"
	"github.com/michaelwp/dragon/handlers"
	"github.com/michaelwp/dragon/middlewares"
	"log"
)

func main()  {
	r := dragon.NewRouter()
	r.Use(middlewares.Logging())
	api := r.Group("/api/v1")
	api.GET("/home", handlers.Home)
	home := api.Group("/home")
	home.POST("/create", handlers.Create)

	log.Fatal(r.Run(":8090"))
}

