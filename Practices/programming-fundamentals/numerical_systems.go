// Write a program that prints a number in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	number := 56                                  // declare a var using short declaration
	fmt.Printf("number in deciaml: %d\n", number) // number in decimal form
	fmt.Printf("number in binary: %b\n", number)  // number in binary representation
	fmt.Printf("number in hex: %#x\n", number)    // number in hexadecimal representation
}
