//Hands-on exercise #3
//Create TYPED and UNTYPED constants. Print the values of the constants.
//solution: https://play.golang.org/p/NutvJXWUx2
package main

import (
	"fmt"
)

const (
	a     = 3
	b int = 4
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}
