package main

import "fmt"

type persons struct {
	first  string
	second string
	age    int
}

type superAgents struct {
	person persons
	ltk    bool
}

type human interface {
	speak()
}

func bar_human(h human) {
	switch h.(type) {
	case persons:
		fmt.Println("I'm Person ", h.(persons).first)
	case superAgents:
		fmt.Println("I'm Super Agent ", h.(superAgents).person.first)

	}
}

func main() {
	p1 := superAgents{
		person: persons{
			first:  "Hryhorii",
			second: "Kravchenko",
			age:    35,
		},
		ltk: true,
	}

	p2 := superAgents{
		person: persons{
			first:  "Lesya",
			second: "Maslii",
			age:    33,
		},
		ltk: false,
	}

	p3 := persons{
		first:  "Dr.",
		second: "Yes",
	}

	fmt.Println(p3)
	p1.speak()
	p2.speak()
	p3.speak()
	bar_human(p1)
	bar_human(p2)
	bar_human(p3)
}

func (s superAgents) speak() {
	fmt.Println("I am ", s.person.first, s.person.second, " - The super agent speak")
}

func (s persons) speak() {
	fmt.Println("I am ", s.first, s.second, " - The person speak")
}
