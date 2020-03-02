package main

import "fmt"

func main() {
	v := 41
	if v == 40 {
		fmt.Println("our value is 40")
	} else if v == 41 {
		fmt.Println("our value is 41")
	} else {
		fmt.Printf("our value is not 40")
	}
}
