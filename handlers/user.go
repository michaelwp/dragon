package handlers

import (
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Body)
}

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user list")
}