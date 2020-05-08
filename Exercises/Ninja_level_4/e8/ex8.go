package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	for _, v := range keys {
		for i, x := range m[v] {
			fmt.Printf("index position: %v \t the value is %v \n", i, x)
		}
	}
	m[`Hryhorii`] = []string{`Wife`, `cat`, `house`}
	for k, v := range m {
		fmt.Printf("The key of map is %v\n", k)
		for p, vv := range v {
			fmt.Printf("index position: %v \t the value is %v \n", p, vv)
		}
	}
}
