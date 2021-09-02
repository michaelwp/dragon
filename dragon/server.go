package dragon

import (
	"log"
	"net/http"
)

func NewRouter() Router {
	return &router{}
}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// validating http method requested
	err := r.isMethodAllowed(rw, req)
	if err != nil {
		return
	}

	// running middleware
	for _, m := range srvMiddlewares {
		err = m(rw, req)
		if err != nil {
			return
		}
	}

	r.Handler(rw, req)
}

func (r *router) Run(address string) error {
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}
