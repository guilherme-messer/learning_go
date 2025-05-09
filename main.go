package main

import (
	"learning_go/exercises"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go exercises.ForLoopExercise1(&wg)
	go exercises.ForLoopExercise2(&wg)
	go exercises.ForLoopExercise3(&wg)

	wg.Wait()
}
