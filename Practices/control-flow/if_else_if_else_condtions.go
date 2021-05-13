// create a program that uses “else if” and “else”.
// take a number check whether it by divisible by 2, 4, 9
// print 2 if the number is divisble by 2
// print 4 if the number is divisble by 4
// print not divisble by 2 and 4 if the number is not divisble by 2 and 4

package main

import "fmt"

func main() {
	number := 45
	if number%2 == 0 {
		fmt.Printf("Number: %v is divisble by 2\n", number)
	} else if number%4 == 0 {
		fmt.Printf("Number: %v is not divisble by 2, but\n", number)
		fmt.Printf("Number: %v is divisble by 4\n", number)
	} else {
		fmt.Printf("Number: %v is not divisble by 2 and 4\n", number)
	}

	//The if statement supports a composite syntax where the condition expression is preceded by an initialization statement.

	if a := 10; a < 0 { // variable "a" declared in this statement are available in all branches
		fmt.Println(a, "is negative")
	} else if a < 10 {
		fmt.Println(a, "contains single digit")
	} else {
		fmt.Println(a, "contains double digits")
	}
}
