package main

import "fmt"

func main() {
	server := NewServer(":3000")
	fmt.Println("Server run in port 3000")
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logger()))
	server.Handle("POST", "/home", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.Listen()
}
