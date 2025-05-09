package exercises

import (
	"fmt"
	"sync"
)

func ForLoopExercise3(wg *sync.WaitGroup) {
	defer wg.Done()

	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Print(arr)
	fmt.Println()
}
