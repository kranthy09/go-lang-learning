package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// create a speak() method for person
func (p person) speak() {
	fmt.Println("Hi, i am", p.first, p.last)
	fmt.Println("I am,", p.age, "years old")
}

// create a struct of type dog
type dog struct {
	breed  string
	origin string
}

// we can create methods with same name
func (a dog) speak() {
	fmt.Println("woohoo!")
}

func main() {

	// create a person with type struct person
	person_1 := person{
		first: "nobitha",
		last:  "doreman",
		age:   10,
	}

	// call the speak method
	person_1.speak()

	// create a dog with type struct dog
	dog_1 := dog{
		breed:  "douborman",
		origin: "Germany",
	}

	// call the speak method defined for dog struct
	dog_1.speak()
}
