package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	h := t.Hour()

	if h < 12 {
		fmt.Println("Good Morning")
	} else if h < 18 {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good Evening")
	}
}
