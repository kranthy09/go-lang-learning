package main

import "fmt"

type person struct {
	first                      string
	last                       string
	favorite_ice_cream_flvaour []string
}

func main() {
	// create a struct of TYPE person
	p1 := person{
		first:                      "kranthi",
		last:                       "kumar",
		favorite_ice_cream_flvaour: []string{"chocolate", "vanilla", "butterscotch"},
	}

	// create a struct of TYPE person
	p2 := person{
		first:                      "nobita",
		last:                       "doremon",
		favorite_ice_cream_flvaour: []string{"blackcurrent", "chocolate", "mango"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v.first)
		for index, value := range v.favorite_ice_cream_flvaour {
			fmt.Println(index, value)
		}
		fmt.Println("-------------------")
	}
}
