package main

import (
	"fmt"
)

func main() {
	number := 42
	fmt.Printf("Binary: "+"%b\n", number)
	fmt.Printf("Decimal: "+"%d\n", number)
	fmt.Printf("Hex: "+"%#x\n", number)
}
