package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	init := time.Now()
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://amazon.com",
	}

	for _, s := range servers {
		reviewServer(s)
	}
	rangeTime := time.Since(init)
	fmt.Println("Pasaron: ", rangeTime)
}

func reviewServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "No esta disponible")
	} else {
		fmt.Println(server, "Funcionando")
	}
}
