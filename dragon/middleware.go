package dragon

import "net/http"

var srvMiddlewares = make([]func(rw http.ResponseWriter, req *http.Request) error, 0)

func (r *router) Use(middleware ...func(rw http.ResponseWriter, req *http.Request) error) {
	srvMiddlewares = append(srvMiddlewares, middleware...)
}
