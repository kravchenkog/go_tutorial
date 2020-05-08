package main

import "fmt"

func main() {
	var mySlice []int
	for i := 42; i <= 51; i++ {
		mySlice = append(mySlice, i)
	}
	fmt.Println(mySlice[:5])
	fmt.Println(mySlice[5:10])
	fmt.Println(mySlice[2:7])
	fmt.Println(mySlice[1:6])
}
