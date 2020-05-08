package main

import "fmt"

func main() {
	m := map[string]int{
		"kasper": 33,
		"gogo":   44,
	}
	m["koko"] = 33
	fmt.Println(m)
	for k, v := range m {
		fmt.Println("k :", k, "   v :", v)
	}
}
