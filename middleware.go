package dragon

import (
	"strings"
)

type (
	middleware struct {
		RouterGroup string
		Middlewares []HandlerFunc
	}
)

var srvMiddlewares = make([]HandlerFunc, 0)
var middlewares = make([]middleware, 0)

func (r *router) Use(m ...HandlerFunc) {
	srvMiddlewares = append(srvMiddlewares, m...)
	middlewares = append(middlewares, middleware{
		RouterGroup: r.RouterGroup,
		Middlewares: srvMiddlewares,
	})
}

func middlewareChecking(d *Dragon) error {
	for _, m := range middlewares {

		// is router group registered with middleware
		isContains := strings.Contains(d.Path, m.RouterGroup)

		// if false then skip process
		if !isContains {
			continue
		}

		// if true then run the middlewares
		for _, f := range m.Middlewares {
			err := f(d)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Dragon) Next() error {
	return nil
}
