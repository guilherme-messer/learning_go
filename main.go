package main

import (
	"fmt"
)

func myfunc() {
	i := 0
Here:
	fmt.Println(i)
	i++
	goto Here
}

func main() {
	myfunc()
}
