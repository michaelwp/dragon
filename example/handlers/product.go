package handlers

import (
	"net/http"

	"github.com/michaelwp/dragon"
)

type (
	product struct {
		ProductId   uint   `json:"product_id"`
		ProductType uint   `json:"product_type"`
		ProductName string `json:"product_name"`
	}

	UdateProductNameParameter struct {
		ProductId   uint   `json:"product_id"`
		ProductName string `json:"product_name"`
	}

	response struct {
		Code    int         `json:"code" xml:"code"`
		Message string      `json:"message" xml:"message"`
		Data    interface{} `json:"data" xml:"data"`
	}
)

func (r response) NewResponse(code int, message string, data interface{}) response {
	r.Code = code
	r.Message = message
	r.Data = data
	return r
}

func (r response) NewErrorResponse(message string) response {
	r.Code = 0
	r.Message = message
	r.Data = nil
	return r
}

func CreateProduct(d *dragon.Dragon) error {
	var productBody product
	var resp response

	err := d.Body(&productBody)
	if err != nil {
		return d.ResponseJSON(http.StatusBadRequest, resp.NewErrorResponse(err.Error()))
	}
	return d.ResponseJSON(http.StatusOK, resp.NewResponse(1, "Create Product", productBody))
}

func ListProduct(d *dragon.Dragon) error {
	var resp response
	return d.ResponseJSON(http.StatusOK, resp.NewResponse(1, "List Of Product", d.Query))
}

func GetProduct(d *dragon.Dragon) error {
	var resp response
	return d.ResponseJSON(http.StatusOK, resp.NewResponse(1, "Get Specific Product", d.Params))
}

func UpdateProductName(d *dragon.Dragon) error {
	var req UdateProductNameParameter
	var resp response

	if err := d.Body(&req); err != nil {
		return d.ResponseJSON(http.StatusBadRequest, resp.NewErrorResponse(err.Error()))
	}

	return d.ResponseJSON(http.StatusOK, resp.NewResponse(1, "success update name", req))
}
