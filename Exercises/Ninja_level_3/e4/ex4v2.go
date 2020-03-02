package main

import "fmt"

func main() {
	y := 1984
	for {
		if y > 2020 {
			break
		} else {
			fmt.Printf("Year %v.  I have been alive\n", y)
		}
		y++
	}

}
