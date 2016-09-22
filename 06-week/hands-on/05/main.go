package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func main() {
	p1 := person{
		"Jeremy",
		"Perkins",
		[]string{"watermelon", "hamburger", "hotdog"},
	}

	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for i, v := range p1.favFood {
		fmt.Println(i, v)
	}
}
