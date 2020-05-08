package main

import "fmt"

func main() {
	p1 := struct {
		firstName  string
		secondName string
		age        int
	}{
		firstName:  "Hryhorii",
		secondName: "Kravchenko",
		age:        35,
	}
	fmt.Println(p1)
}
