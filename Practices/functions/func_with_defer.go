package main

import "fmt"

func rockStar() {
	fmt.Println("called rock star function")
}

func yatch() {
	fmt.Println("Rockstar hagouts in yatch")
}

func past() {
	fmt.Println("Rockstar has past, but it has to be proved")
}

func main() {
	defer yatch() // defer function calls when all the functions around it has retuned or called
	rockStar()    // first, calls rockstar function
	past()        //second, calls past function, later yatch will be called
}
