package main

import "fmt"

type gator int

func main() {

	var g1 gator
	g1 = 27

	var x int

	x = int(g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
