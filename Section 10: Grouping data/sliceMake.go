package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[8] = 76
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	//x[10] = 33 ERROR
	x = append(x, 23)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 23)
	x = append(x, 231)
	x = append(x, 232)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
