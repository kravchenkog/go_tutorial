package main

import "fmt"

type person struct {
	first  string
	second string
}

type superAgent struct {
	person person
	ltk    bool
}

func (s superAgent) speak() {
	fmt.Println("I am ", s.person.first, s.person.second)
	fmt.Println("and I have the ltk: ", s.ltk)
}

func main() {
	sa1 := superAgent{
		person: person{
			first:  "Hryhorii",
			second: "Kravchenko",
		},
		ltk: true,
	}
	sa2 := superAgent{
		person: person{
			first:  "Lesya",
			second: "Maslii",
		},
		ltk: false,
	}
	sa1.speak()
	sa2.speak()
}
