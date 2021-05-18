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

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("Favorite ice cream flavours of", p1.first, p1.last)
	for index, value := range p1.favorite_ice_cream_flvaour {
		fmt.Println(index, value)
	}

	fmt.Println("Favorite ice cream flavours of", p2.first, p2.last)
	for index, value := range p2.favorite_ice_cream_flvaour {
		fmt.Println(index, value)
	}

	// creating struct of anonymous type
	anonymous_struct := struct {
		name string
		age  int
	}{
		name: "skywalker",
		age:  25,
	}

	fmt.Println("print anonymous struct")
	fmt.Println(anonymous_struct)

}
