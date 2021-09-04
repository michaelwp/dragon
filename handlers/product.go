package handlers

import (
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create product")
}

func ListProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "product list")
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get product")
}