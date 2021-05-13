// 1. Using the following operators, write expressions and assign their values to variables
// 2. print the variables

package main

import (
	"fmt"
)

func main() {
	// create variables using short declaration
	a := 45
	b := 67

	//applying operator expressions
	isEqualTo := (a == b)
	fmt.Printf("%v == %v: %v\n", a, b, isEqualTo)

	isLessthanOrEqualto := (a <= b)
	fmt.Printf("%v <= %v: %v\n", a, b, isLessthanOrEqualto)

	isGreaterthanOrEqualto := (a >= b)
	fmt.Printf("%v >= %v: %v\n", a, b, isGreaterthanOrEqualto)

	isNotEqualto := (a != b)
	fmt.Printf("%v != %v: %v\n", a, b, isNotEqualto)

	isLessthan := (a < b)
	fmt.Printf("%v < %v: %v\n", a, b, isLessthan)

	isGreaterthan := (a > b)
	fmt.Printf("%v > %v: %v\n", a, b, isGreaterthan)
}
