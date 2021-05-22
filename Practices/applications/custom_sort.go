package main

import (
	"fmt"
	"sort"
)

// create type user
type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// implementing sort.Interface for []user based on Age field
type ByAge []user

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// implementing sort.Interface for []user based on Last field
type ByLast []user

func (a ByLast) Len() int {
	return len(a)
}
func (a ByLast) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByLast) Less(i, j int) bool {
	return a[i].Last < a[j].Last
}

// implementing sort.Interfaaec for []string based on user.Sayings field
type BySayingForEachUser []string

func (a BySayingForEachUser) Len() int {
	return len(a)
}
func (a BySayingForEachUser) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a BySayingForEachUser) Less(i, j int) bool {
	return a[i] < a[j]
}

// create a function to print user in a pleasent way
func printUsers(users []user) {
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println("\t", s)
		}
	}
}

func main() {
	// create users u1, u2, u3
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	// create slice of users: u1, u2, u3
	users := []user{u1, u2, u3}

	fmt.Println("--------users, unsorted: -------")
	printUsers(users)

	// sort by Age
	sort.Sort(ByAge(users)) // sorts the []user by Age
	fmt.Println("-------users, sorted by Age: -------")
	printUsers(users)

	// sort by Last
	sort.Sort(ByLast(users)) // sorts the []user by Last
	fmt.Println("-------users, sorted by Last: -------")
	printUsers(users)

	// sort each user's Sayings
	fmt.Println("-------Each User sayings, sorted-------")
	for _, v := range users {
		sort.Sort(BySayingForEachUser(v.Sayings))
		fmt.Println("user: with", v.First, ",sorted Sayings: ", v.Sayings)
	}

}
