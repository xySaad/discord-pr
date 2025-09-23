package mux

import (
	"io"
	"net/http"
)

type Mux[D any] struct {
	*http.ServeMux
	Dependecy D
}

func New[D any](dependecy D) Mux[D] {
	return Mux[D]{&http.ServeMux{}, dependecy}
}

type Request struct {
	*http.Request
	body []byte
}

func (r *Request) Body() []byte {
	if r.body != nil {
		return r.body
	}

	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		return nil
	}
	defer r.Request.Body.Close()

	r.body = body
	return r.body
}

type Hanlder[D any] func(http.ResponseWriter, *Request, D)
type Middleware[D any] func(http.ResponseWriter, *Request, D) bool

func (mx *Mux[D]) Route(pattern string, hanlder Hanlder[D], middlewares ...Middleware[D]) {
	mx.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		req := &Request{r, nil}
		for _, middleware := range middlewares {
			if !middleware(w, req, mx.Dependecy) {
				return
			}
		}

		hanlder(w, req, mx.Dependecy)
	})
}
