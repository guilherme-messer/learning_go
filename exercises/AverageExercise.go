package exercises

import (
	"fmt"
	"sync"
)

func AverageExercise(slice []float64, wg *sync.WaitGroup) {
	defer wg.Done()

	length := len(slice)

	var sum float64

	for i := 0; i < length; i++ {
		sum += slice[i]
	}

	result := sum / float64(length)

	fmt.Printf("Average Exercise: %.2f\n", result)
}
