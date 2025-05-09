package exercises

import (
	"fmt"
	"sync"
)

func ForLoopExercise1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}
