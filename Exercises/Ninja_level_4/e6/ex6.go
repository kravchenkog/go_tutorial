package main

import "fmt"

func main() {
	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	statesSlice := make([]string, len(states), 500)
	for i := 0; i < len(states); i++ {
		statesSlice[i] = fmt.Sprintf("Position %d", i)
	}
	fmt.Println("The length of my slice is :", len(statesSlice))
	fmt.Println("The capacity of my slice is :", cap(statesSlice))
	fmt.Println("The slice is :", statesSlice, "\n")

	statesSlice = append(statesSlice, states...)
	fmt.Println("The length of my slice is :", len(statesSlice))
	fmt.Println("The capacity of my slice is :", cap(statesSlice))
	fmt.Println("The slice is :", statesSlice, "\n")

	for i := 0; i < len(statesSlice); i++ {
		fmt.Println(i, statesSlice[i])
	}

}
