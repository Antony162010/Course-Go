package main

import "fmt"

type animal interface {
	move() string
}

type dog struct{}

func (dog) move() string {
	return "I am a dog"
}

type cat struct{}

func (cat) move() string {
	return "I am a cat"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	p := dog{}
	moveAnimal(p)
	c := cat{}
	moveAnimal(c)
}
