package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 3, 5, 1, 7}
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:4])
	fmt.Println(x[1:4])
	for i := 0; i < len(x); i++ {
		fmt.Println(i, "\t", x[i])
	}

}
