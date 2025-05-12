package main

import (
	"learning_go/exercises"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	go exercises.ForLoopExercise1(&wg)
	go exercises.ForLoopExercise2(&wg)
	go exercises.ForLoopExercise3(&wg)

	val := []float64{10, 20, 30}
	go exercises.AverageExercise(val, &wg)

	go exercises.FizzBuzz(&wg)

	wg.Wait()
}
