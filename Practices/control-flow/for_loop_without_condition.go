// Create a for loop using this syntax
// .....for { }
// Have it print out the years you have been alive

package main

import "fmt"

func main() {
	dob := 1996
	presentYear := 2021
	for { // if we do not mention the condition, for loop considers it as true always
		fmt.Println("Year i have been Alive", dob)
		dob++
		if dob > presentYear { // condition to break out of loop
			break
		}
	}
}
