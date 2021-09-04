package main

import (
	"github.com/michaelwp/dragon/dragon"
	"github.com/michaelwp/dragon/handlers"
	"log"
)

func main() {
	r := dragon.NewRouter()
	api := r.Group("/api/v1")

	// user
	user := api.Group("/users")
	user.GET("", handlers.ListUser)
	user.POST("", handlers.CreateUser)
	user.GET("/1", handlers.GetUser)

	// product
	product := api.Group("/products")
	product.GET("", handlers.ListProduct)
	product.POST("", handlers.CreateProduct)
	product.GET("/1", handlers.GetProduct)

	// home
	home := api.Group("/home")
	home.GET("", handlers.Home)

	log.Fatal(r.Run(":8090"))
}
