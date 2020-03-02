package main

import "fmt"

func main() {
	var ar [5]int
	fmt.Println(ar)
	ar[4] = 2
	fmt.Println(ar, "\t", len(ar))
}
