package main

import "fmt"

func main() {
	favSport := "soccer"
	switch {
	case favSport == "ping-pong":
		fmt.Println("Yes its ping-pong")
	case favSport == "basketball":
		fmt.Println("Yes its basketball")
	case favSport == "soccer":
		fmt.Println("Yes its soccer")

	}

}
