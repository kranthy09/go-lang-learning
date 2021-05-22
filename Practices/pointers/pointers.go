package main

import "fmt"

// create type person struct
type person struct {
	first string
	last  string
	age   int
}

// create func which takes person as pointer and changes the first and last names for a person
func changeMe(p *person, first string, last string) {
	p.first = first
	p.last = last
}

func main() {
	// create a person
	p1 := person{
		first: "nobita",
		last:  "doreman",
		age:   10,
	}
	fmt.Println(p1) // print the person p1

	// pass the pointer of p1
	changeMe(&p1, "kranthi", "kumar")
	fmt.Println(p1) // print the new values of person p1
}
