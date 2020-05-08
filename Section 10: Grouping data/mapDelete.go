package main

import "fmt"

func main() {
	m := map[string]int{
		"Hryhorii": 35,
		"Lesya":    33,
	}
	delete(m, "Hryhorii")
	fmt.Println(m)
	m["Hryhorii"] = 35
	fmt.Println(m)
	delete(m, "Hryhorii2")
	fmt.Println(m)
}
