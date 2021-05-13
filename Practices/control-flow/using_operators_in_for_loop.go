// Print out the remainder (modulus)
// which is found for each number between 10 and 100 when it is divided by 4.

package main

import "fmt"

func main() {
	// declare variables and assign 10, 100 values using short declaration
	a := 10
	b := 100

	// using for loop, prints rem for numbers between 10, 100
	for i := a; i <= b; i++ {
		fmt.Printf("Remainder: %v, when divided %v by %v\n", i%4, i, 4)
	}
}
