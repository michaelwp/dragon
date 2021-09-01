package main

import (
	"errors"
	"log"
	"net/http"
	"reflect"
)

type Router interface {
	GET(address string, handlerFunc http.HandlerFunc)
	POST(address string, handlerFunc http.HandlerFunc)
	Run(address string) error
	Group(address string) *router
}

type router struct {
	RouterGroup string
	Path        string
	Handler     http.HandlerFunc
	Method      string
}

func NewRouter() Router {
	return &router{}
}

func (r *router) Run(address string) error {
	log.Printf("Server listen and serve on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r *router) Group(path string) *router {
	r.RouterGroup = r.RouterGroup + path
	return r
}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := isMatching(r.Method, req.Method)
	if err != nil {
		rw.Write([]byte("Method " + err.Error()))
		return
	}

	r.Handler(rw, req)
}

func isMatching(expect string, actual string) error {
	if !reflect.DeepEqual(expect, actual) {
		return errors.New("not match")
	}

	return nil
}

/* METHODS */

func (r router) GET(path string, handlerFunc http.HandlerFunc) {
	r.Method = http.MethodGet
	r.Handler = handlerFunc
	r.Path = r.RouterGroup + path

	http.Handle(r.Path, &r)
}

func (r router) POST(path string, handlerFunc http.HandlerFunc) {
	r.Method = http.MethodPost
	r.Handler = handlerFunc
	r.Path = r.RouterGroup + path

	http.Handle(r.Path, &r)
}
