package main

import "fmt"

func main() {
	valores := []int{10, 20, 30}
	fmt.Println(mapFunc(dobra, valores))
}

func mapFunc(f func(int) int, valores []int) []int {
	resultado := make([]int, len(valores))

	for i, x := range valores {
		resultado[i] = f(x)
	}

	return resultado
}

func dobra(x int) int {
	return x * 2
}
