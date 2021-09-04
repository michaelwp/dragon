package dragon

import "net/http"

type router struct {
	RouterGroup string
	Path        string
	Handler     []handler
}

type handler struct {
	Handler     http.HandlerFunc
	Methods     string
}

type middleware struct {
	RouterGroup string
	Middlewares []func(rw http.ResponseWriter, req *http.Request) error
}
