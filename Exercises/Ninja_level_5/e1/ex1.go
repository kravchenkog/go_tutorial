//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import "fmt"

type person struct {
	firstName  string
	secondName string
	iceCream   string
}

func main() {
	p1 := person{
		firstName:  "Alesya",
		secondName: "Maslii",
		iceCream:   "chcolate",
	}
	p2 := person{
		firstName:  "Hryhorii",
		secondName: "Kravchenko",
		iceCream:   "vanila",
	}
	fmt.Println(p1.iceCream, p2.iceCream)
}
