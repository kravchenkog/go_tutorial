package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 3, 5, 1, 7}
	x = append(x, 4, 5, 6)
	fmt.Println(x)
	y := []int{55, 66, 77, 88}
	x = append(x, y...) // !!!
	fmt.Println(x)

}
