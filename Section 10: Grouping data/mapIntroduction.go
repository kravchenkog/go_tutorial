package main

import "fmt"

func main() {
	m := map[string]int{
		"Hryhorii Kravhenko": 663475037,
		"Masliy Lesya":       959147759,
	}
	fmt.Println(m)
	fmt.Println("print by key: ", m["Hryhorii Kravhenko"])
	fmt.Println("print by invalid key: ", m["Putin"])
	v, ok := m["Putin"]
	fmt.Println("v : ", v)
	fmt.Println("ok : ", ok)

	if _, ok := m["Putin"]; ok == false {
		fmt.Println("Putin is not exist")
	}

}
