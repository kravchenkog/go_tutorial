package main

import (
	"fmt"
)

const (
	c = iota
	e = 2019 + iota
	f = 2019 + iota
	g = 2019 + iota
	w = 2019 + iota
)

func main() {

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(w)
}
