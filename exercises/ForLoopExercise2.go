package exercises

import (
	"fmt"
	"sync"
)

func ForLoopExercise2(wg *sync.WaitGroup) {
	defer wg.Done()

	i := 0
Loop:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Loop
	}

	fmt.Println()
}
