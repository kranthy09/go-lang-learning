// Create a program that uses a switch statement with the switch expression specified as a variable

package main

import (
	"fmt"
)

func main() {
	// declare an int type variable using short declaration
	number := 2
	switch number { // expression evaluates and looks for each case expression and executes that case
	case 1:
		fmt.Println(number, "is the first natural number")
	case 2:
		fmt.Println(number, "is second natural number") // prints this statement as number is 2.
	case 3:
		fmt.Println(number, "is third natural number")
	default: // default case executes if any of the above case is not true
		fmt.Println(number, "is a natural number greater than 3")
	}
}
