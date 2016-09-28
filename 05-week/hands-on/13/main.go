package main

import "fmt"

type gator int

func main() {

	var g1 gator
	g1 = 27

	var x int

	x = g1

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//not possible to assign g1 to x because they are two different types
	//main.go:14: cannot use g1 (type gator) as type int in assignment
}