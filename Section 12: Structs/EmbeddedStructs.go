package main

import "fmt"

func main() {
	type person struct {
		firstName  string
		secondName string
		age        int
	}
	type secretAgent struct {
		person
		ltk bool
	}

	p1 := person{
		firstName:  "Hryhorii",
		secondName: "Kravchenko",
		age:        35,
	}
	p2 := person{
		firstName:  "Maslii",
		secondName: "Lesia",
		age:        33,
	}
	agent := secretAgent{
		person: person{
			firstName:  "James",
			secondName: "Bond",
		},
		ltk: true,
	}

	fmt.Println(p1.firstName, p1.secondName, p1.age)
	fmt.Println(p2.firstName, p2.secondName, p2.age)
	fmt.Println(agent.firstName, agent.person.secondName, agent.age, agent.ltk)
}
