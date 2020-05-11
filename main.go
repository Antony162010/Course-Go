package main

import "fmt"

func main() {
	fmt.Println("Hello world.")
	imprime()
}

func imprime() {
	defer fmt.Println("Inicio")
	fmt.Println("Final")
}
