package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// CheckAuth is ...
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking authentication")
			if flag {
				f(rw, r)
			} else {
				return
			}
		}
	}
}

// Logger is ...
func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(rw, r)
		}
	}
}
