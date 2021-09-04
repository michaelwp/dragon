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
	hh, err := isPathExist(rw, req)
	if err != nil {
		return
	}

	//// running middleware
	//for i, m := range middlewares {
	//	log.Println(i)
	//	for _, f := range m.Middlewares {
	//		err = f(rw, req)
	//		if err != nil {
	//			return
	//		}
	//	}
	//}

	hh.Handler(rw, req)
}

func (r *router) Run(address string) error {
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}
