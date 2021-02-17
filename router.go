package main

import (
	"net/http"
)

// Router is ...
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// NewRouter is ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// FindHandler is ...
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exists := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exists, methodExist
}

// ServerHTTP is ...
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
