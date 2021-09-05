package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon"
)

func Home(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, "home")
}