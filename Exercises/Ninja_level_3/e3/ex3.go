package main

import "fmt"

func main() {
	bornYear := 1984
	currentYear := 2020
	for bornYear <= currentYear {
		fmt.Printf("Year %v.  I have been alive\n", bornYear)
		bornYear++
	}

}
