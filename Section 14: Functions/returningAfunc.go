package main

import "fmt"

func main() {

	foo8()(38)
	fmt.Printf("%T \n", foo8())
	foo9(55)
}

func foo8() func(int) {
	b := func(n int) {
		fmt.Println("My name is Hryhorii, my age is ", n)
	}
	return b
}

func foo9(a int) {
	func(a int) {
		fmt.Println("The A is ", a)
	}(a)

}
