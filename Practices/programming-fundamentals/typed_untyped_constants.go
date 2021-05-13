// Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import (
	"fmt"
)

const (
	// declare the untyped constant
	a = 34
	// declare typed constant with type "int"
	b int = 32
)

func main() {
	fmt.Printf("untyped constant a of value: %v and of type: %T\n", a, a)
	fmt.Printf("typed constant b of value: %v and of type: %T\n", b, b)
}
