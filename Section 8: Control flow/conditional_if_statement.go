package main

import "fmt"

func main() {
	if x := 42; x == 22 {
		fmt.Println("True")
	}
	//fmt.Println(x) = error because x not in scope

}
