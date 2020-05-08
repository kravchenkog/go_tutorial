package main

import "fmt"

func main() {
	x := foo1(1, 3, 4, 5, 3, 2, 4, 5, 6, 3, 234, 234)
	fmt.Println(x)
}

func foo1(x ...int) int {
	fmt.Printf("%T \n", x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
