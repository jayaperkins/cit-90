package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Todd"] = 47
	m["Jeremy"] = 20
	m["Patrick"] = 27

	fmt.Println("map: ", m)
}
