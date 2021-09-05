package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon/dragon"
)

func Home(d *dragon.Dragon) {
	fmt.Fprint(d.ResponseWriter, "home")
}