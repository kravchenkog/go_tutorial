package main

import (
	"fmt"
)

//func (receiver) identifier(parameters) (returns) { code }

func main() {
	foo()
	s := "Hello Hryhorii"
	bar(s)
	fmt.Println("5*5= ", barbara(5))
	d, f := multy("Hryhorii", "Alesya")
	fmt.Println(d)
	fmt.Println(f)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println(s)
}

func barbara(v int) int {
	return v * v
}

func multy(string2 string, string3 string) (string, string) {
	s2 := fmt.Sprint("Maslii ", string3)
	s1 := fmt.Sprint("Kravchenko ", string2)
	return s1, s2
}
