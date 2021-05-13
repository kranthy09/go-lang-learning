// 1. use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
// 2. then use the “=” operator to ASSIGN that value to “y”
// 3. print out the value stored in “y”
// 4. print out the type of “y”

package main

import (
	"fmt"
)

// create "myType" which has underlying type "int"
type myType int

// create variable of myType
var x myType = 54

// create variable of type int
var y int

func main() {
	y = int(x)                                                            // conversion from myType to int and assign it to var y
	fmt.Println(y)                                                        //prints 42
	fmt.Printf(
		"var:%v has type: %T\n var:%v has type: %T\n", x, x, y, y
	) // prints the type of the x, y
}
