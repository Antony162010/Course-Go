package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	init := time.Now()
	canal := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
		"http://amazon.com",
	}

	i := 0
	for {
		if i > 10 {
			break
		}
		for _, s := range servers {
			go reviewServer(s, canal)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		i++
	}

	// for i := 0; i < len(servers); i++ {
	// 	fmt.Println(<-canal)
	// }
	rangeTime := time.Since(init)
	fmt.Println("Pasaron: ", rangeTime)
}

func reviewServer(server string, canal chan string) {
	_, err := http.Get(server)
	if err != nil {
		// fmt.Println(server, "No esta disponible")
		canal <- server + " no disponible"
	} else {
		// fmt.Println(server, "Funcionando")
		canal <- server + " funcionando"
	}
}
