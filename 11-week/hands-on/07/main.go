package main

import "fmt"

func main() {

	ii := [3]int{14, 15, 42}
	for idx, val := range ii {
		fmt.Println(idx, val)
	}

}
