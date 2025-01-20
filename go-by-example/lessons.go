package main

import (
	"fmt"
	"time"
)

func variables() {
	// typing variables
	var a string = "hello world"
	var b, c int = 1, 2
	var d bool = true

	// default values
	var e int    // 0
	var f string // ""
	var g bool   // false

	// short-hand variale assignment (needs to be in a function)
	h := "apple"

	fmt.Println(a, b, c, d, e, f, g, h)
}

func loops() {
	// single condition
	i := 1
	for i <= 3 {
		fmt.Println("single condition", i)
		i = i + 1
	}

	// initialize, condition, after
	for j := 4; j <= 6; j++ {
		fmt.Println("three sisters", j)
	}

	// for range
	for k := range 3 {
		fmt.Println("range", k)
	}

	// for break
	for {
		fmt.Println("loop")
		break
	}

	// continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Odd only", n)
	}

}

func runSwitch() {
	// basic
	i := 2
	switch i {
	case 1:
		fmt.Println("Basic i = 1 ")
	case 2:
		fmt.Println("Basic i = 2")
	case 3:
		fmt.Println("Basic i = 3 ")
	}

	// multiple statements in a case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend!")
	default:
		fmt.Println("Its the weekday")
	}

	// no prop in switch
	j := "apple"
	switch {
	case j == "orange":
		fmt.Println("I'm an orange!")
	case j == "apple":
		fmt.Println("I'm an apple")
	default:
		fmt.Println("no fruits found here")
	}

	// type finder
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("whatAmI? I'm a bool!", t)
		case int:
			fmt.Println("whatAmI? I'm a int!", t)
		case string:
			fmt.Println("whatAmI? I'm a string!", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}

// arrays have a specific length
func runArrays() {

}

// slices are dynamic arrays (automatically readjust size)
func runSlice() {
	// Uninitlaize slice is nil and has length 0
	var s []string
	fmt.Println("uninitialized slice:", s, "Length of slice:", len(s) == 0, "Slice is nil : ", s == nil)

	// Initialize an empty slice with make
	s = make([]string, 5)
	fmt.Println("initialized slice:", s, "Length of slice:", len(s) == 0, "Slice is nil : ", s == nil, "Slice capacity : ", cap(s))

	// set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Add items via index s[0] = `a` etc.")
	fmt.Println("initialized slice:", s, ", Length of slice:", len(s) == 0, ", Slice is nil : ", s == nil, ", Slice capacity : ", cap(s))

	// get
	fmt.Println("Item at index 2 is :", s[2])

	// append
	s = append(s, "e")
	s = append(s, "f")
	s = append(s, "g")

	fmt.Println("initialized slice:", s, ", Length of slice:", len(s) == 0, ", Slice is nil : ", s == nil, ", Slice capacity : ", cap(s))

	// DEEP copy a whole array
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy of slice: ", c)

	c[0] = "hello"
	fmt.Println("Copy of slice modified: ", c)
	fmt.Println("Original of slice: ", s)

}

func runMaps() {
	// delcare a map
	myMap := make(map[string]int)

	// set key, value
	myMap["key1"] = 1
	myMap["key2"] = 2
	myMap["key3"] = 3
	fmt.Println("myMap:", myMap)
	value, present := myMap["key1"]
	fmt.Println("value, present: ", value, present)

	// get(key) = value
	fmt.Println("myMap[key1]:", myMap["key1"])

	// delete(key)
	delete(myMap, "key1")

	// clear(map)
	clear(myMap)

	// use the second return value to indicate if the key is present
	a, b := myMap["key10"]
	fmt.Println("a, b", a, b)

	// declare and intialized with values
	fruitMap := map[string]int{"apple": 1, "blueberry": 2, "cherry": 3}
	fmt.Println(fruitMap)
}

func runAddFunc(a, b, c int) int {
	return a + b + c
}

func runMultipleReturnValues(a, b int) (int, int) {
	return a + 1, b + 1

	// a, b := runMultipleReturnValues(1, 2)
	// fmt.Println(a, b)

	// _, c := runMultipleReturnValues(3, 4)
	// fmt.Println(c)
}

func runVariadicFunc(nums ...int) {

	if nums == nil {
		fmt.Println("nums is nil")
	}

	for i, num := range nums {
		fmt.Printf("i=%d, num=%d\n", i, num)
	}

	// runVariadicFunc([]int{1, 2, 3}...)
}
