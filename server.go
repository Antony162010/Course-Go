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

// Handle is ...
func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
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
