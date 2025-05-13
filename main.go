package main

import (
	"fmt"
	"learning_go/stack" // ajuste para o seu módulo se for diferente
	"strconv"
	"strings"
)

func main() {
	expr := "* + 2 3 4"
	resultado := polishCalculator(expr)
	fmt.Printf("Resultado de \"%s\" = %d\n", expr, resultado)
}

func polishCalculator(s string) int {
	tokens := strings.Split(s, " ")
	pilha := stack.New()

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if num, err := strconv.Atoi(token); err == nil {
			pilha.Push(num)
		} else {
			b := pilha.Pop()
			a := pilha.Pop()
			var resultado int

			switch token {
			case "+":
				resultado = a + b
			case "-":
				resultado = a - b
			case "*":
				resultado = a * b
			case "/":
				if b == 0 {
					panic("Divisão por zero")
				}
				resultado = a / b
			default:
				panic("Operador inválido: " + token)
			}

			pilha.Push(resultado)
		}
	}

	return pilha.Pop()
}
