package main

import (
	"encoding/json"
	"net/http"
)

// Middleware is ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// MetaData is ...
type MetaData interface{}

// User is ...
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// ToJSON is ...
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}
