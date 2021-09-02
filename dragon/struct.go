package dragon

import "net/http"

type router struct {
	RouterGroup string
	Path        string
	Handler     http.HandlerFunc
	Method      string
}
