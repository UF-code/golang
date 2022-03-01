package main

import "fmt"

func main() {

	value_1 := 10

	fmt.Printf("Value: %v", value_1)
	// Value: 10
	fmt.Printf("\nMemory Location of '%v': %v", value_1, &value_1)
	// Memory Location of '10': 0xc00001a098

	value_2 := "Go is Fun!"

	fmt.Printf("\nValue: %v", value_2)
	// Value: Go is Fun!
	fmt.Printf("\nMemory Location of '%v': %v", value_2, &value_2)
	// Memory Location of 'Go is Fun!': 0xc000044230

}
