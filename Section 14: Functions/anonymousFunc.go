package main

import (
	"fmt"
)

func main() {
	func(a int) {
		fmt.Println("Anonimus func ")
		fmt.Println(a)
	}(123)
	foo6()
}

func foo6() {
	fmt.Println("foo6 is running")
}
