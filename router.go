package main

import (
	"net/http"
)

// Router is ...
type Router struct {
	rules map[string]http.HandlerFunc
}

// NewRouter is ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

// FindHandler is ...
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exists := r.rules[path]
	return handler, exists
}

// ServerHTTP is ...
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists := r.FindHandler(request.URL.Path)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, request)
}
