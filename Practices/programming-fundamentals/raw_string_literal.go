// Create a variable of type string using a raw string literal. Print it.

package main

import (
	"fmt"
)

func main() {
	// raw string can be created using back ticks(``)
	// raw strings will be ommited as they are, when declared
	raw_string := `this is raw string
	this line appears in new line
	this is third line
	"Hello, World"
	I am doing great in golang`
	fmt.Println(raw_string)
}
