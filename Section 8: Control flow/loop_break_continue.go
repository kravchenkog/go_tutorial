package main

import "fmt"

func main() {
	s := 0
	for {
		s++
		if s > 100 {
			break
		}
		if s%2 != 0 {
			continue
		}

		fmt.Println(s)
	}

}
