package exercises

import (
	"fmt"
	"strconv"
	"sync"
)

func FizzBuzz(wg *sync.WaitGroup) {
	defer wg.Done()

	var valores [100]string

	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			valores[i-1] = "FizzBuzz"
			continue
		} else if i%3 == 0 {
			valores[i-1] = "Fizz"
			continue
		} else if i%5 == 0 {
			valores[i-1] = "Buzz"
			continue
		} else {
			valores[i-1] = strconv.Itoa(i)
		}
	}

	fmt.Println(valores)
}
