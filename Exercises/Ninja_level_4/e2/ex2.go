package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("1. create a SLICE of TYPE int")
	x := []int{}
	fmt.Println(x)

	fmt.Println("2. assign 10 VALUES ")
	for i := 0; i <= 10; i++ {
		x = append(x, i+rand.Int())
	}
	fmt.Println(x)

	fmt.Println("3. Range over the slice and print the values out.")
	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Println("4. Using format printin print out the TYPE of the slice")
	fmt.Printf("The format of variable is %T", x)

}
