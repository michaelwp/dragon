package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon"
	"net/http"
)

type user struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"user_name"`
}

func CreateUser(d *dragon.Dragon) error {
	var userBody user
	var resp response

	err := d.Body(&userBody)
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
		resp.Data = nil

		return d.ResponseXML(http.StatusBadRequest, resp)
	}

	resp.Code = 1
	resp.Message = "Create User"
	resp.Data = userBody

	return d.ResponseXML(http.StatusOK, resp)
}

func ListUser(d *dragon.Dragon) error {
	var resp response

	resp.Code = 1
	resp.Message = fmt.Sprintf("List Of %v User", d.Query.Get("gender"))
	resp.Data = map[string]interface{}{
		"token":   d.GetContext("token"),
		"user_id": d.GetContext("userID"),
	}

	return d.ResponseJSON(http.StatusOK, resp)
}

func GetUser(d *dragon.Dragon) error {
	var resp response

	resp.Code = 1
	resp.Message = fmt.Sprintf("Get User %v", d.Params["id"])
	resp.Data = map[string]interface{}{
		"token":   d.GetContext("token"),
		"user_id": d.GetContext("userID"),
	}

	return d.ResponseJSON(http.StatusOK, resp)
}
