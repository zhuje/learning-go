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

func runRange() {
	// for range loop to sum a slice of numbers
	// for range to find a value in a slice of number

	arr := []int{1, 2, 3}
	for i, value := range arr {
		fmt.Println(i, value)
		if value == 3 {
			fmt.Println("Found the value 3 at index: ", i)
		}
	}

	// for range to print out the key value pairs of map[string]string
	fruits := map[string]string{"apple": "amy", "blueberry": "bob"}
	for key, val := range fruits {
		fmt.Println(key, val)
	}

}

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {

	fmt.Println("zeroptr(ptr *int), ptr = ", ptr)
	*ptr = 0 // access the value
}

func runPointers() {
	i := 1
	zeroval(i)
	fmt.Println("Pass value:", i)

	zeroptr(&i) // send the address
	fmt.Println("Pass pointer:", i)
}

func main() {

	runPointers()

}
