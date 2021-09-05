package handlers

import (
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

		return d.ResponseJSON(http.StatusBadRequest, resp)
	}

	resp.Code = 1
	resp.Message = "Create User"
	resp.Data = userBody

	return d.ResponseJSON(http.StatusOK, resp)
}

func ListUser(d *dragon.Dragon) error {
	var resp response

	resp.Code = 1
	resp.Message = "List Of User"
	resp.Data = d.Query.Get("gender")

	return d.ResponseJSON(http.StatusOK, resp)
}

func GetUser(d *dragon.Dragon) error {
	var resp response

	resp.Code = 1
	resp.Message = "Get Specific User"
	resp.Data = d.Params["id"]

	return d.ResponseJSON(http.StatusOK, resp)
}
