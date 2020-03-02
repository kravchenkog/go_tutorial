package main

import "fmt"

type chevrolet int

var x chevrolet = 33

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
