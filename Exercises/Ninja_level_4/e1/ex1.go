package main

import "fmt"

func main() {
	fmt.Println("create an ARRAY which holds 5 VALUES of TYPE int")
	x := make([]int, 5, 100)
	fmt.Println(x, "\n")

	fmt.Println("assign VALUES to each index position \n")
	for v, _ := range x {
		x[v] = v
	}
	fmt.Println(x)

	println("Range over the array and print the values out. \n ")
	for v, _ := range x {
		fmt.Println(x[v])
	}

	println("print out the TYPE of the array. \n ")
	fmt.Printf("the type is %T", x)

}
