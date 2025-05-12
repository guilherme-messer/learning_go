package main

import "fmt"

func main() {
	stack := []rune{10, 20, 30}
	var top rune

	stack = push(stack, 40)
	fmt.Println(stack)

	stack, top = pop(stack)
	fmt.Println(stack, top)
}

func push(stack []rune, valor rune) []rune {
	stack = append(stack, valor)

	return stack
}

func pop(stack []rune) ([]rune, rune) {
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return stack, top
}
