// Using the short declaration operator,
//ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 42
// “James Bond”
// true
// Now print the values stored in those variables using
// a single print statement
// multiple print statements

package main

import "fmt"

func main() {
	number := 42
	name := "James Bond"
	boolValue := true

	//a single print statement
	fmt.Println("**print using single print statement**")
	fmt.Println(number, name, boolValue)

	//multiple print statement
	fmt.Println("**print using multiple print statements**")
	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(boolValue)
}
