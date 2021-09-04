package dragon

import (
	"net/http"
	"strings"
)

var srvMiddlewares = make([]func(rw http.ResponseWriter, req *http.Request) error, 0)
var middlewares = make([]middleware, 0)

func (r *router) Use(m ...func(rw http.ResponseWriter, req *http.Request) error) {
	srvMiddlewares = append(srvMiddlewares, m...)
	middlewares = append(middlewares, middleware{
		RouterGroup: r.RouterGroup,
		Middlewares: srvMiddlewares,
	})
}

func middlewareChecking(rw http.ResponseWriter, req *http.Request) error {
	for _, m := range middlewares {

		// is router group registered with middleware
		isContains := strings.Contains(req.URL.Path, m.RouterGroup)

		// if false then skip process
		if !isContains {
			continue
		}

		// if true then run the middlewares
		for _, f := range m.Middlewares {
			err := f(rw, req)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
