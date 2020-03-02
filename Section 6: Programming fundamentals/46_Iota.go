package main

import (
	"fmt"
)

const (
	e = iota
	f
	g
)
const (
	w = iota
	r
	t
	y
)

func main() {

	fmt.Println(w)
	fmt.Println(r)
	fmt.Println(t)
	fmt.Println(y)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", t)
}
