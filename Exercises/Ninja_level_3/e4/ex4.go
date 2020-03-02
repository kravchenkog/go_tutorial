package main

import "fmt"

func main() {
	bornYear := 1984
	currentYear := 2020
	for i := bornYear; i <= currentYear; i++ {
		fmt.Printf("Year %v.  I have been alive\n", i)
	}

}
