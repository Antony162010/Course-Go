package main

import "net/http"

// Middleware is ...
type Middleware func(http.HandlerFunc) http.HandlerFunc
