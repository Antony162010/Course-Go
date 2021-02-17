package main

import (
	"encoding/json"
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

// PostRequest is ...
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var meta MetaData
	err := decoder.Decode(&meta)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	fmt.Fprintf(w, "Payload %v\n", meta)
}

// PostUser is ...
func PostUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	response, err := user.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "Payload %v\n", user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
