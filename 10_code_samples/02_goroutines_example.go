package main

import (
	"fmt"
	"time"
)

func display(word string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\n%v %v", word, i)
	}
}

func main() {
	// Calling Goroutine
	go display("Function with Goroutine")

	// Calling normal function
	display("Function without Goroutine")

	// Function without Goroutine 0
	// Function with Goroutine 0
	// Function with Goroutine 1
	// Function without Goroutine 1
	// Function without Goroutine 2
	// Function with Goroutine 2
	// Function with Goroutine 3
	// Function without Goroutine 3
	// Function without Goroutine 4
}
