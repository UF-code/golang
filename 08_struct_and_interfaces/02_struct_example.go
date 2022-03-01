package main

import "fmt"

func main() {
	person := struct {
		first      string
		last       string
		age        int
		founder_of string
	}{

		first:      "Anatoly",
		last:       "Yakovenko",
		age:        44,
		founder_of: "Solana",
	}

	full_name := person.first + " " + person.last
	fmt.Printf("Full Name: %v", full_name)
	// Full Name: Anatoly Yakovenko
	fmt.Printf("\nAge: %v", person.age)
	// Age: 44
	fmt.Printf("\nFounder Of: %v", person.founder_of)
	// Founder Of: Solana
}
