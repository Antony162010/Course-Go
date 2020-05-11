package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) makeOperation(input string, operator string) int {
	valores := strings.Split(input, operator)

	operador1 := parseText(valores[0])
	operador2 := parseText(valores[1])
	switch operator {
	case "+":
		return operador1 + operador2
	case "-":
		return operador1 - operador2
	case "*":
		return operador1 * operador2
	case "/":
		return operador1 / operador2
	default:
		return 0
	}
}

func parseText(input string) int {
	operador, err1 := strconv.Atoi(input)

	if err1 != nil {
		fmt.Println(err1)
	}
	return operador
}

func readInput() string {
	scaner := bufio.NewScanner(os.Stdin)
	scaner.Scan()

	return scaner.Text()
}
func main() {
	operations := []string{"+", "-", "*", "/"}

	fmt.Print("Ingresa tu suma: ")
	operation := readInput()

	operator := ""
	for i := 0; i < len(operations); i++ {
		if strings.Contains(operation, operations[i]) {
			operator = operations[i]
		}
	}

	c := calc{}
	result := c.makeOperation(operation, operator)
	fmt.Print("Resultado: ", result)

}
