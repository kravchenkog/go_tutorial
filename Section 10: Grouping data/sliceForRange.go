package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 6}
	fmt.Println(x)
	fmt.Println(x[len(x)-1]-x[len(x)-2], "\n\n")
	for i, v := range x[1:] {
		fmt.Printf("i : %v \t, v: %v, \n", i, v)
	}
}
