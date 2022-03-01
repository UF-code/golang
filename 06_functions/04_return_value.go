package main

import "fmt"

func main() {
	fmt.Println("I am", name()) // I am James Bond

	name, age := name_and_age()
	fmt.Println("I am", name, "and I am", age, "years old")
	// I am James Bond and I am 42 years old
}
func name() string {
	name := "James Bond"
	return name
}
func name_and_age() (string, int) {
	name := name()
	age := 42

	return name, age
}
