package main

import "fmt"

func main() {
	defer foo4() //In the end of the method
	bar4()
}

func foo4() {
	fmt.Println("f003")
}

func bar4() {
	fmt.Println("bar4")
}
