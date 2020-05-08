package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 77, 54, 3, 33, 3, 2, 3, 3, 4, 4, 4, 3}
	x := foo2(xi...)
	fmt.Println(x)
	y := foo3("Hr")
	fmt.Println(y)
}

func foo2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func foo3(s string, x ...int) string {
	sum := 0
	for _, v := range x {
		sum += v
	}
	d := fmt.Sprint(sum, s)

	return d
}
