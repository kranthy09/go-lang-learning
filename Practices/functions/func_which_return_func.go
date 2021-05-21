package main

import (
	"fmt"
)

// create a func which returns func
func watchMe(name string) func(author string) string {
	intro := `My name is: ` + name + ".\n"
	return func(author string) string {
		dialogue := intro + `I have listen watchMe song. ` + `Author of song is: ` + author
		return dialogue
	}
}
func main() {
	// assing function to var
	f := watchMe("nobitha")

	// call the returned function by passing required arguments
	result := f("silento")

	// print the result
	fmt.Println(result)
}
