package main

import (
	"fmt"
	"math"
)

// interface for geometry (area() and perimeter())
type geometry interface {
	area() float64
	perimeter() float64
}

// struct for rectangle
type rectangle struct {
	height float64
	width  float64
}

// struct for circle
type circle struct {
	radius float64
}

type square struct {
	width float64
}

// implement interface geometry: function for area (rectangle)
func (r rectangle) area() float64 {
	return r.height * r.width
}

// implement interface geometry: function for permetrier (rectangle)
func (r rectangle) perimeter() float64 {
	return (2 * r.height) * (2 * r.width)
}

// implement interface geometry: function for area (circle)
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// implement interface geometry: function for perimeter (circle)
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s square) area() float64 {
	return s.width * s.width
}

func (s square) perimeter() float64 {
	return 4 * s.width
}

// type assertion
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("cirlce with radius", c.radius)
	} else {
		fmt.Println("not a circle")
	}
	fmt.Println("--------------------------")
}

func detectShape(g geometry) {
	switch shape := g.(type) {
	case circle:
		fmt.Println("circle with radius", shape.radius)
	case rectangle:
		fmt.Println("rectangle with width", shape.width, "and height", shape.height)
	case square:
		fmt.Println("square with with", shape.width)
	default:
		fmt.Println("unknown shape")
	}
	fmt.Println("--------------------------")
}

func measure(g geometry) {
	fmt.Println("geometry struct: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimeter: ", g.perimeter())
}

func TestInterface() {
	r := rectangle{
		width:  3,
		height: 4,
	}
	c := circle{
		radius: 5,
	}
	s := square{
		width: 6,
	}

	measure(r)
	detectCircle(r)
	detectShape(r)

	measure(c)
	detectCircle(c)
	detectShape(c)

	measure(s)
	detectShape(s)

}
