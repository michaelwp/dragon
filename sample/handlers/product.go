package handlers

import (
	"github.com/michaelwp/dragon"
	"net/http"
)

type (
	product struct {
		ProductId   uint   `json:"product_id"`
		ProductType uint   `json:"product_type"`
		ProductName string `json:"product_name"`
	}

	response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    interface{}
	}
)

func CreateProduct(d *dragon.Dragon) error {
	var productBody product
	var resp response

	err := d.Body(&productBody)
	if err != nil {
		resp.Code = 0
		resp.Message = err.Error()
		resp.Data = nil

		return d.ResponseJSON(http.StatusBadRequest, resp)
	}

	resp.Code = 1
	resp.Message = "Create Product"
	resp.Data = productBody

	return d.ResponseJSON(http.StatusOK, resp)
}

func ListProduct(d *dragon.Dragon) error {
	var resp response

	resp.Code = 1
	resp.Message = "List Of Product"
	resp.Data = d.Query

	return d.ResponseJSON(http.StatusOK, resp)
}

func GetProduct(d *dragon.Dragon) error {
	var resp response

	resp.Code = 1
	resp.Message = "Get Specific Product"
	resp.Data = d.Params

	return d.ResponseJSON(http.StatusOK, resp)
}
