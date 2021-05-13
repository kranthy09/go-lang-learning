// Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	switch { // switch case with no switch expression, always executes the case with true expression
	case true:
		fmt.Println("prints always true, as we did not mentioned the switch expression")
	case false:
		fmt.Println("it is false")
	}
}
