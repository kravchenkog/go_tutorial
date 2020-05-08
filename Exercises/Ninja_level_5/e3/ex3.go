//Create a new type: vehicle.
//The underlying type is a struct.
//The fields:
//doors
//color

//Create two new types: truck & sedan.
//The underlying type of each of these new types is a struct.
//Embed the “vehicle” type in both truck & sedan.
//Give truck the field “fourWheel” which will be set to bool.
//Give sedan the field “luxury” which will be set to bool. solution

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	fourWheel bool
	vehicle
}
type sedan struct {
	luxury bool
	vehicle
}

func main() {
	s1 := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
	}
	s2 := truck{
		fourWheel: false,
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
	}
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(s1.vehicle.color, s1.vehicle.doors, s1.luxury)

}
