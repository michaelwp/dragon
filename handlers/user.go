package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/michaelwp/dragon/dragon"
)

type user struct {
	Name string `json:"name"`
}

func CreateUser(d *dragon.Dragon) {
	var userBody user

	json.NewDecoder(d.Body).Decode(&userBody)
	fmt.Fprint(d.ResponseWriter, userBody)
}

func ListUser(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, d.Query.Get("gender"))
}

func GetUser(d *dragon.Dragon) {
	content := fmt.Sprintf("User ID: %s", d.Params["id"])
	fmt.Fprint(d.ResponseWriter, content)
}
