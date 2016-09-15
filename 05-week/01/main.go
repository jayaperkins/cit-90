package main

import (
	"math"
	"fmt"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea (x shape) {
	fmt.Println("My area is:", x.area())
}

func main () {
	s1 := square{34}
	c1 := circle{27}
	printArea(s1)
	printArea(c1)
}
