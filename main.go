package main

import (
	"fmt"
)

func main() {
	variableArgs(10, 20, 30, 40, 50, 60, 90, 70)
}

func variableArgs(args ...int) {
	for _, i := range args {
		fmt.Printf("%d\n", i)
	}
}
