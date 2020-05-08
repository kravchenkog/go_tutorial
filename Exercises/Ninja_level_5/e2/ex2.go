//Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
//solution: https://play.golang.org/p/RvvLyAMvGo

package main

import "fmt"

type person struct {
	firstName  string
	secondName string
	favorite   []string
}

func main() {
	p1 := person{
		firstName:  "Lesya",
		secondName: "Maslii",
		favorite: []string{
			"chocolate", "martini", "orange",
		},
	}
	p2 := person{
		firstName:  "Hryhorii",
		secondName: "Kravchenko",
		favorite:   []string{"vodka", "beer", "coca-cola"},
	}

	persons := map[string]person{
		p1.secondName: p1,
		p2.secondName: p2,
	}

	for k, v := range persons {
		fmt.Println("SecondName :", k, "\n", "FirstName :", v.firstName)
		for y, b := range v.favorite {
			fmt.Println(y+1, b)
		}
	}
}
