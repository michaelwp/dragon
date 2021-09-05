package dragon

import (
	"context"
	"io"
	"net/http"
)

type (
	router struct {
		RouterGroup string
		Path        string
		Handler     []handler
		Params      []*param
	}

	param struct {
		Key      string
		Position int
		Value    string
	}

	handler struct {
		Handler HandlerFunc
		Methods string
	}

	middleware struct {
		RouterGroup string
		Middlewares []func(rw http.ResponseWriter, req *http.Request) error
	}

	Dragon struct {
		ResponseWriter http.ResponseWriter
		Request        *http.Request
		Context        context.Context
		Params         map[string]string
		Headers        http.Header
		Path           string
		Body           io.ReadCloser
		Queries        map[string]string
		RemoteAddress  string
	}

	HandlerFunc func(d *Dragon)
)
