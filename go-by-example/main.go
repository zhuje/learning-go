package main

import (
	"fmt"
)

/*
*

	Anonymous functions (func() int{}) is used to create a closure.
	The closure function runClosure() will keep i in the memory stack
	for each call to the anonymous function.
*/
func runClosure() func() int {
	i := 0

	return func() int {
		i++
		return i
	}

	// incrementClosure := runClosure()

	// fmt.Println(incrementClosure())
	// fmt.Println(incrementClosure())
}

func main() {

}
