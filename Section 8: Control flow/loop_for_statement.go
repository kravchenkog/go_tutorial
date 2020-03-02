package main

import "fmt"

func main() {
	x := 0
	for x <= 10 && x != 11 {
		fmt.Printf("x is : %d\n", x)
		x += 1
	}
	y := 0
	for {
		if y != 10 {
			fmt.Printf("Y != 10 ; Y is [%d]\n", y)
			y += 1
		} else {
			break
		}

	}
}
