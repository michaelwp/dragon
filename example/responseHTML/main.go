package main

import (
	"github.com/michaelwp/dragon"
	"log"
)

func main() {
	r := dragon.NewRouter()

	r.GET("/home", func(d *dragon.Dragon) error {
		return d.ResponseHTML(
			200,
			"<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Dragon</title>\n</head>\n<body>\n    <h1>DRAGON</h1>\n</body>\n</html>")
	})

	log.Fatal(r.Run(":8090"))
}
