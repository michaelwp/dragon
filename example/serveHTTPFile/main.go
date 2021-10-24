package main

import (
	"github.com/michaelwp/dragon"
	"log"
)

func main() {
	r := dragon.NewRouter()

	// serving static http file
	r.ServeHTTPFile("./static")

	log.Fatal(r.Run(":8090"))
}
