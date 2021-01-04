package main

import "net/http"

// Server is ...
type Server struct {
	port   string
	router *Router
}

// NewServer is ...
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Listen is ...
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port, nil)
	if error != nil {
		return error
	}
	return nil
}
