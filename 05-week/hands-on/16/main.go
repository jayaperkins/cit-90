package main

import "fmt"

type gator int

type flamingo bool

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func main() {

	var g gator

	var f flamingo

	bayou(g)

	bayou(f)
}

