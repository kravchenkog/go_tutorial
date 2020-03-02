package main

import (
	"fmt"
)

const (
	a int = 2
	b     = 2.2
	c     = "Test"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
