package main

import "fmt"

func main() {
	ii := []int{2, 3, 4, 5, 6, 7}
	s := sum(ii...)
	fmt.Println(s)

	s2 := even(sum, ii...)
	fmt.Println(s2)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var fi []int
	for _, v := range vi {
		if v%2 == 0 {
			fi = append(fi, v)
		}

	}
	s := f(fi...)
	return s
}
