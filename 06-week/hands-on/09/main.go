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

func (t truck) transportationDevice() string {
	return fmt.Sprintln("This truck has", t.doors, "doors, and", t.fourWheel, "has four wheel drive." )
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("This sedan has", s.doors, "doors, and", s.luxury, "is very comfortable" )
}

func main() {
	v1 := truck{
		vehicle{
			2,
			"Blue",
		},

		true,
	}

	v2 := sedan{
		vehicle{
			4,
			"Red",
		},

		false,
	}

	fmt.Println(v2)

	fmt.Println(v1.fourWheel)

	fmt.Println(v1.transportationDevice())

	fmt.Println(v2.transportationDevice())
}
