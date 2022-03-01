package main

import "fmt"

func main() {

	value := 10
	address_of_value := &value

	fmt.Printf("Memory Location: %v", address_of_value)
	// Memory Location: 0xc0000aa058
	fmt.Printf("\nValue of Memory Location '%v': %v", address_of_value, *address_of_value)
	// Value of Memory Location '0xc0000aa058': 10
}
