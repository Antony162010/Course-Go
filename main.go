package main

import "fmt"

func main() {
	server := NewServer(":3000")
	fmt.Print("Server run in port 3000")
	server.Listen()
}
