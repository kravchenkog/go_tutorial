package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	x = append(x, 2)
	fmt.Println(x)

}
