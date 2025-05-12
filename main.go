package main

import (
	"fmt"
	"time"
)

func main() {
	fibonacci()
}

func fibonacci() {
	start := time.Now()

	fibonacci0(1, 1)

	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução: %s\n", elapsed)
}

func fibonacci0(x1 int, x2 int) {
	fmt.Printf("%d\n", x1)

	x1, x2 = x2, x1+x2

	if x1 > 2000000000 {
		return
	}

	fibonacci0(x1, x2)
}
