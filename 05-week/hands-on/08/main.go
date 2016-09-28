package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle{
			2,
			"Blue",
		},

		true,
	}

	v2 := truck{
		vehicle{
			4,
			"Red",
		},

		false,
	}

	fmt.Println(v2)

	fmt.Println(v1.fourWheel)

}
