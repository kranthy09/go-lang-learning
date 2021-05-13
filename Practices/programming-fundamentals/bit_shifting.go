// 1. Write a program that,
// 2. assigns an int to a variable
// 3. prints that int in decimal, binary, and hex
// 4. shifts the bits of that int over 1 position to the left, and assigns that to a variable
// 5. prints that variable in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	// create a variable using short declaration
	x := 45
	fmt.Printf("number in decimal representation: %d\nnumber in binary representation: %b\nnumber in hex representation: %#x\n", x, x, x)

	// bit shifting over 1 position towards left and assigning the value to another variable
	y := x << 1
	fmt.Printf("number after shifting over 1 position towards left: %v\n", y)
	fmt.Printf("new number in decimal representation: %d\nnew number in binary representation: %b\nnew number in hex representation: %#x\n", y, y, y)
}
