//Hands-on exercise #4
//Write a program that
//assigns an int to a variable
//prints that int in decimal, binary, and hex
//shifts the bits of that int over 1 position to the left, and assigns that to a variable
//prints that variable in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	variable := 44
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %x\n", variable, variable, variable)
	variableNew := variable << 1
	fmt.Printf("Decimal: %d, Binary: %b, Hex: %x", variableNew, variableNew, variableNew)
}
