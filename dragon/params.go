package dragon

import (
	"strings"
)

func registerParamsUrl(path string) []*param {
	var (
		url    = strings.Split(path, "/")
		params []*param
	)

	for i, u := range url {
		if key, isTrue := isParams([]byte(u)); isTrue {
			var p param

			// form param struct
			p.Key = key
			p.Position = i
			p.Value = ""

			// register param struct to params slice
			params = append(params, &p)
		}
	}

	return params
}

func isParams(char []byte) (string, bool) {
	if len(char) > 0 && char[0] == ':' {
		return string(char[1:]), true
	}
	return "", false
}

func (r *router) formParamsUrl(path string) string {
	var (
		rrPath   = strings.Split(r.Path, "/")
		reqPath  = strings.Split(path, "/")
		joinPath string
	)

	if len(rrPath) == len(reqPath) {

		// iterate over the params registered
		// and assign the param value
		for _, p := range r.Params {
			rrPath[p.Position] = reqPath[p.Position]
			p.Value = reqPath[p.Position]
		}

		// is the path matching
		if isMatching(rrPath, reqPath) {
			joinPath = strings.Join(rrPath, "/")
			return joinPath
		}
	}

	return r.Path
}

func (r *router) mapParamsUrl() map[string]string {
	var mParam = map[string]string{}

	for _, p := range r.Params {
		mParam[p.Key] = p.Value
	}

	return mParam
}
