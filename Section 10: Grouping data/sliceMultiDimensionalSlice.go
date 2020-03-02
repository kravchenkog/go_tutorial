package main

import "fmt"

func main() {
	jb := []string{"Hryhorii", "Kravchenko", "35", "man"}
	fmt.Println(jb)
	ml := []string{"Lesya", "Masliy", "33", "women"}
	fmt.Println(ml)
	xp := [][]string{jb, ml}
	fmt.Println(xp)

}
