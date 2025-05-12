package main

import "fmt"

func main() {
	valores := []int{10, 20, 30, 60, 50, 15}

	fmt.Println(bubbleSort(valores))
}

func bubbleSort(valores []int) []int {
	n := len(valores)
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if valores[i-1] > valores[i] {
				valores[i-1], valores[i] = valores[i], valores[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return valores
}
