package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	mux         *mux.Router
	middlewares []func(http.Handler) http.Handler
}

func NewRouter() *Router {
	return &Router{
		mux:         mux.NewRouter(),
		middlewares: make([]func(http.Handler) http.Handler, 0),
	}
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc(pattern, handler)
}

// Methods adds a matcher for HTTP methods
func (r *Router) Methods(methods ...string) *mux.Route {
	return r.mux.Methods(methods...)
}

func (r *Router) Use(middleware func(http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handler http.Handler = r.mux
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i](handler)
	}
	handler.ServeHTTP(w, req)
}
