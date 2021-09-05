package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon/dragon"
)

func CreateProduct(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, "create product")
}

func ListProduct(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, "product list")
}

func GetProduct(d *dragon.Dragon) {
	content := fmt.Sprintf("Product Type: %s, name: %s", d.Params["type"], d.Params["name"])
	fmt.Fprint(d.ResponseWriter, content)
}