package main

import "fmt"

func main() {
	for i := 10; i < 100; i++ {
		i++
		h := i % 4
		fmt.Printf("Value [%#v] remainder [4] = %v \n", i, h)
	}

}
