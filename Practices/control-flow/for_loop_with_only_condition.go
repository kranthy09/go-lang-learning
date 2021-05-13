// Create a for loop using this syntax
//.....for condition { }
// Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	dob := 1996
	presentYear := 2021
	for dob <= presentYear { // loop breaks when the condition fails
		fmt.Println("Year i have been alive", dob) //prints each year which i have alive till 2021
		dob++
	}
}
