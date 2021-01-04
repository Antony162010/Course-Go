package main

import (
	"fmt"
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

// ServerHTTP is ...
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello")
}
