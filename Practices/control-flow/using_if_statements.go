// Create a program that shows the “if statement” in action.

package main

import "fmt"

func main() {
	// declare a string type variable
	name := "kranthi"
	if name == "kranthi" { // if block executes only if the codition is true
		fmt.Println(`Logged in as "kranthi"`)
	}

	// declare a int type variable
	number := 54
	if number%2 == 0 { // if block executes only if the codition is true
		fmt.Printf("Number:%v, is a Prime number\n", number)
	}
}
