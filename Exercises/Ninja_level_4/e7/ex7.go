package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i := 0; i < len(x); i++ {
		for y := 0; y < len(x[i]); y++ {
			fmt.Println(x[i][y])
		}
	}
	fmt.Println("*****************************************")
	for _, v := range x {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
