package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	var guess int

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite um número:")
		fmt.Scanf("%d", &guess)
		if guess == num {
			fmt.Println("Acertou")
			break
		}
		fmt.Println("Errou feio")
	}
	fmt.Println(num)
}
