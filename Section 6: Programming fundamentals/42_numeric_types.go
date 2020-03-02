package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 42.5344534534
	a = a + 1
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
