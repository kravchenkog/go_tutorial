package main

import "fmt"

type person struct {
	firstName, secondName string
	age                   int32
}

func main() {
	p1 := person{
		firstName:  "Hryhorii",
		secondName: "Kravchenko",
		age:        35,
	}
	p2 := person{
		firstName:  "Lesya",
		secondName: "Maslii",
		age:        33,
	}

	fmt.Println(p1.age, p1.secondName, p1.firstName)
	fmt.Println(p2.age, p2.secondName, p2.firstName)
}
