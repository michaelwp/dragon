package handlers

import (
	"fmt"
	"github.com/michaelwp/dragon"
)

func Home(d *dragon.Dragon) error {
	fmt.Fprint(d.ResponseWriter, "home")
	return nil
}