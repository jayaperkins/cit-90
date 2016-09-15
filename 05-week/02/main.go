package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

type members interface {
	speak() string
}

func (p person) speak() string {
	return fmt.Sprintln(p.fName)
}

func (sa secretAgent) speak() string {
	return fmt.Sprintln(sa.fName)
}

func people(x members) {
	fmt.Println("Hello:", x.speak())
}

func main() {
	p := person{"Jeremy", "Perkins", 20}
	s := secretAgent{person{"Zachary", "Perkins", 27}, false}

	fmt.Println(p.fName)
	fmt.Println(s.fName)

	people(p)
	people(s)
}
