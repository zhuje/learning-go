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

func runArrays() {

}

func main() {
	runSwitch()
}
