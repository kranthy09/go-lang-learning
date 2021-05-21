package main

import "fmt"

func foo() int {
	fmt.Println("Function foo is called")
	return 43
}

func bar() (string, int) {
	fmt.Println("Function bar is called")
	return "bar function is called and returned", 66
}

func main() {
	x := foo()
	fmt.Println(x)
	str, num := bar()
	fmt.Println(str, num)
}
