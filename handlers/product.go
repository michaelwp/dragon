package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon/dragon"
)

type product struct {
	ProductId   uint   `json:"product_id"`
	ProductType uint   `json:"product_type"`
	ProductName string `json:"product_name"`
}

func CreateProduct(d *dragon.Dragon) {
	var productBody product

	err := d.Body(&productBody)
	if err != nil {
		fmt.Fprint(d.ResponseWriter, err)
		return
	}

	content := fmt.Sprintf(
		"product_id: %v \nproduct_type: %v \nproduct_name: %v",
		productBody.ProductId,
		productBody.ProductType,
		productBody.ProductName,
	)

	fmt.Fprint(d.ResponseWriter, content)
}

func ListProduct(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, "product list")
}

func GetProduct(d *dragon.Dragon) {
	content := fmt.Sprintf(
		"Product Type: %s, name: %s",
		d.Params["type"],
		d.Params["name"],
	)

	fmt.Fprint(d.ResponseWriter, content)
}
