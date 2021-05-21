package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	// assign a function to variable
	f := greet
	f("nobita")

	// create anonymous function and assign to a variable
	g := func(age int) {
		fmt.Println("anonymous function is called")
		fmt.Println("My age respective this function is: ", age)
	}
	g(10)

	// we can assign to another varaible and call that function
	a := g
	fmt.Println("with respective to 'a' variable")
	a(24)
}
