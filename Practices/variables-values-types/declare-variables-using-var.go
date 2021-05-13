package main

import (
	"fmt"
)

// when we declare variables using "var" keyword
//variables will be having package level scope
var x int
var y string
var z bool

// compiler assigns zero values to each identifier when we do not assign values to the identifier

func main() {
	fmt.Println(x) // for int type compiler assigns 0
	fmt.Println(y) // for string type compiler assigns ""
	fmt.Println(z) // for bool type compiler assigns false
}
