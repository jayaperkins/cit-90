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

type transportation interface {
	transportationDevice() string
}

func report (g transportation) {
	fmt.Println(g.transportationDevice())
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

	report(v1)

	report(v2)

}
