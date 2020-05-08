package main

import "fmt"

func main() {
	p1 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Hryhorii",
		lastName:  "Kravchenko",
	}
	fmt.Println(p1)
}
