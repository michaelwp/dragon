package handlers

import (
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "create user")
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "user list")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get user")
}