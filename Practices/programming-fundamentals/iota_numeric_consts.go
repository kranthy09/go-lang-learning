// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import (
	"fmt"
)

// iota makes it easy to declare sequentially growing numeric constants in
const (
	presentYear     = 2021 + iota
	afterOneYear    = 2021 + iota
	afterTwoYears   = 2021 + iota
	afterThreeYears = 2021 + iota
)

// iota sets the value of Low to 0
// and instructs the compiler that the following constants have increasing numeric values.

// Every time a keyword Const appears in the code, iota’s increment is reset.
const (
	a = iota
	b = iota
)

func main() {
	fmt.Printf("present year: %v\n", presentYear)          // prints 2021
	fmt.Printf("after one year: %v\n", afterOneYear)       // prints 2022
	fmt.Printf("after two years: %v\n", afterTwoYears)     // prints 2023
	fmt.Printf("after three years: %v\n", afterThreeYears) // prints 2024

	// as the we have declared a, b constants with iota, iota value will be reset to zero
	fmt.Printf("value of a: %v\n", a) // prints 0
	fmt.Printf("value of a: %v\n", b) // print 1
}
