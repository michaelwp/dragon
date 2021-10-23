package main

import (
	"log"

	"github.com/michaelwp/dragon"
	"github.com/michaelwp/dragon/example/handlers"
)

func main() {
	r := dragon.NewRouter()

	// api
	api := r.Group("/api/v1")

	// home
	home := api.Group("/home")
	home.GET("", handlers.Home)

	// user
	user := api.Group("/users")
	user.Use(Authorize())
	user.GET("", handlers.ListUser)
	user.POST("", handlers.CreateUser)
	user.GET("/:id", handlers.GetUser)

	// product
	product := api.Group("/products")
	product.GET("", handlers.ListProduct)
	product.POST("", handlers.CreateProduct)
	product.GET("/:type/:name", handlers.GetProduct)
	product.PATCH("", handlers.UpdateProductName)

	log.Fatal(r.Run(":8090"))
}
