package main

import (
	"fmt"
	"learning_go/stack"
)

func main() {
	pilha := stack.New()
	pilha.Push(50)
	pilha.Push(100)
	pilha.Pop()
	fmt.Println(pilha)
}
