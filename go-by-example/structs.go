package main

import "fmt"

type person struct {
	name string
	age  int
}

/*
*person, *int // types of pointers
&p // creates the pointer
*p // dereferences the pointer to get the value at this address

returns a pointer to person (*person)
return a point to int (*int)
*/
func newPerson(name string, age int) (*person, *int) {
	i := 1
	p := person{name, age}

	// & create pointers from these values
	return &p, &i
}

func runStruct() {
	p, i := newPerson("abby", 25)
	// address
	fmt.Println(p)
	fmt.Println(i)

	// value (dereferenced)
	fmt.Println(*p)
	fmt.Println(*i)
}

type Cat struct {
	name  string
	age   int
	fuzzy bool
}

func createCatPtr(name string, age int, fuzzy bool) *Cat {
	newCat := Cat{name, age, fuzzy}
	return &newCat
}

func createCat(name string, age int, fuzzy bool) Cat {
	newCat := Cat{name, age, fuzzy}
	return newCat
}

func runCatStruct() {
	bobCatPtr := createCatPtr("bob", 1, true)
	cathyCat := createCat("cathy", 9, false)

	fmt.Println(*bobCatPtr)
	bobCatPtr.age = 4
	fmt.Println(*bobCatPtr)

	fmt.Println(cathyCat)
	bobCatPtr.age = 4
	fmt.Println(cathyCat)
}
