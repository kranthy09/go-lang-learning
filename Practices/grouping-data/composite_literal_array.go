//1. Using a COMPOSITE LITERAL:
//		a. create an ARRAY which holds 5 VALUES of TYPE int
//		b. assign VALUES to each index position.
//2. Range over the array and print the values out.
//3. Using format printing
//		a. print out the TYPE of the array

package main

import "fmt"

func main() {

	// create an array (COMPOSITE LITERAL)
	arr := [4]int{2, 3, 5, 6}
	fmt.Println(arr)

	// range over the array
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// print out TYPE of the array
	fmt.Printf("TYPE of arr: %T\n", arr)

	// create a sclie of TYPE STRING using composite literal
	states := []string{

		` Alabama`, ` Alaska`, ` Arizona`,
		` Arkansas`, ` California`, ` Colorado`,
		` Connecticut`, ` Delaware`, ` Florida`,
	}

	// prints length of slice
	fmt.Println(`length of "state":`, len(states))

	// prints capacity of slice
	fmt.Println(`Capacity of "state"`, cap(states))

	// prints the values, along with index
	for i := 0; i < len(states); i++ {
		fmt.Println("index: ", i, "value: ", states[i])
	}
}
