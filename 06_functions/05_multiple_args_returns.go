package main

import "fmt"

func main() {
	summary, average := summary_and_average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("Summary: %v Average: %v", summary, average) // Summary: 45 Average: 5
}

// Multiple Arguments and Return Values
func summary_and_average(numbers ...int) (int, int) {
	summary := 0

	for _, value := range numbers {
		summary += value
	}

	average := summary / (len(numbers))
	return summary, average
}
