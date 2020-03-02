package main

import (
	"fmt"
)

func main() {
	st := "Hello, playground"
	fmt.Println(st)
	fmt.Printf("%T\n", st)
	li := []byte(st)
	fmt.Println(li)
	for i := 0; i < len(st); i++ {
		fmt.Println("%#U ", st[i])
	}
	for i, v := range st {
		fmt.Println(i, " ", string(v))
	}

}
