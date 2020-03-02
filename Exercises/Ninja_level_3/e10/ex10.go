package main

import "fmt"

func main() {
	fmt.Println(true && true, "\t should be true")
	fmt.Println(true && false, "\t should be false")
	fmt.Println(true || true, "\t should be true")
	fmt.Println(true || false, "\t should be true")
	fmt.Println(!true, "\t should be false")

}
