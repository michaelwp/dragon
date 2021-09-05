package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon/dragon"
)

type user struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

func CreateUser(d *dragon.Dragon) {
	var userBody user

	err := d.Body(&userBody)
	if err != nil {
		fmt.Fprint(d.ResponseWriter, err)
		return
	}

	content := fmt.Sprintf(
		"user_id: %v \nuser_name: %v",
		userBody.UserId,
		userBody.UserName,
	)

	fmt.Fprint(d.ResponseWriter, content)
}

func ListUser(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, d.Query.Get("gender"))
}

func GetUser(d *dragon.Dragon) {
	content := fmt.Sprintf(
		"User ID: %s",
		d.Params["id"],
	)

	fmt.Fprint(d.ResponseWriter, content)
}
