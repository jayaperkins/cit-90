package main

import "fmt"

func main() {
	m := map[string]int{"Jeremy": 20, "Zachary": 27, "Norm": 48}
	fmt.Println(m)

	for k := range m {
		fmt.Println(k)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
