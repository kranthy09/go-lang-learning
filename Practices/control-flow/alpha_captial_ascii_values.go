// Print every rune code point of the uppercase alphabet three times

package main

import "fmt"

func main() {

	// declare variables by assigning ascii values of cap A and cap Z using short declaration
	acii_value_of_cap_A := 65
	ascii_value_of_cap_z := 90

	for i := acii_value_of_cap_A; i <= ascii_value_of_cap_z; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ { // prints each capital value three times
			fmt.Printf("\t%#U\n", i) // prints the rune(character) for int value of each i
		}
	}
}
