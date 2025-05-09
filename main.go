package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for pos, val := range arr {
		fmt.Println(pos, val)
	}

	fmt.Println("----")

	slice := arr[:]
	slice[9] = 99
	fmt.Println(slice[9])
	// slice[10] = 0 // fora de indice

	slice2 := append(slice, 10)
	fmt.Println(slice2)

	fmt.Println("----")

	mes_ano := map[int]string{1: "Janeiro", 2: "Fevereiro"}
	valor := mes_ano[1]

	fmt.Println(valor)
}
