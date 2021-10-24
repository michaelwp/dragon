package handlers

import (
	"github.com/michaelwp/dragon"
	"net/http"
)

func Home(d *dragon.Dragon) error {
	var resp response

	resp.Message = "Hello World"
	return d.ResponseJSON(http.StatusOK, resp)
}