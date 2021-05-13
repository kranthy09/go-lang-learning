//1. Create your own type. Have the underlying type be an int.
//2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
//3. in func main
//		a. print out the value of the variable “x”
//		b. print out the type of the variable “x”
//		c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
//		d. print out the value of the variable “x”

package main

import (
	"fmt"
)

// create my own type, myType" which has underlying type "int"
type myType int

//create a variable of "myType"
var x myType

func main() {
	fmt.Println(x)        //prints zero value
	fmt.Printf("%T\n", x) //prints the type of var x
	x = 42                //assign an int value
	fmt.Println(x)        //prints 42
}
