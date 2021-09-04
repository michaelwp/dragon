package dragon

import (
	"net/http"
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
