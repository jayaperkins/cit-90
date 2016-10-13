package main

import "fmt"

type shape struct {
	name string
}

type circle struct {
	shape
	diameter int
}

type square struct {
	shape
	length int
}

type large interface {
	size() bool

}

func (c circle) size() bool {
	return true
}

func (s square) size() bool {
	return false
}


func main() {

	s1 := circle{
		shape{
			"Circle",
		},
		10,
	}

	s2 := square{
		shape{
			"Square",
		},
		7,
	}

	fmt.Println(s1)
	fmt.Println(s2)

}
