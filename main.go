package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pilha stack
	pilha.push(50)
	pilha.push(100)
	pilha.pop()
	fmt.Printf("stack %v\n", pilha)
}

// Função String
func (s stack) String() string {
	var str string
	for i := 0; i < s.head; i++ {
		str = str + "[" + "Index:" +
			strconv.Itoa(i) + " : " + "Data:" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
