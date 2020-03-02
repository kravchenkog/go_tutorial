//Hands-on exercise #2
//Using the following operators, write expressions and assign their values to variables:
//==
//<=
//>=
//!=
//<
//>

package main

import (
	"fmt"
)

func main() {
	value1 := 43
	value2 := 44

	fmt.Println(value1 == value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 > value2)
	f := 44 == 45
	fmt.Println("44 == 45: ", f)
}
