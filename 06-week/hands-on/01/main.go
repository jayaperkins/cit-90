package main

import (
	"fmt"
)

func main() {

	s := []int{43, 6, 27}
	fmt.Println(s)

	for i := range s {
		fmt.Println(i)
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
