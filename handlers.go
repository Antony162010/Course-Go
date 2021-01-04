package main

import (
	"fmt"
	"net/http"
)

// HandleRoot is ...
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

// HandleHome is ...
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello home")
}
