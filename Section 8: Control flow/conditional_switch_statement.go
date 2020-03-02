package main

import "fmt"

func main() {
	i := 2
	switch b := i; {
	case b == 3, b == 2: //OR STATEMENT
		fmt.Println("it's can be true, because b == 2")

	case b == 1:

		fmt.Println("it's false")

	case b == 2:

		fmt.Println("it's true")
		fallthrough

	case b != 3:

		fmt.Println("it's true, b != 3")
		fallthrough
	case b != 4:

		fmt.Println("it's true, i != 4")
		fallthrough

	default:

		fmt.Println("default case")

	}
}
