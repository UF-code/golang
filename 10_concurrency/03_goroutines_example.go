package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func worker(word string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("Worker %v goroutine\n", word)
	}
	defer wg.Done()
}

func main() {
	fmt.Printf("Total CPU: %v\n", runtime.NumCPU())
	fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())

	wg.Add(1)
	go worker("1")

	wg.Add(1)
	go worker("2")

	defer fmt.Printf("Goroutines: %v\n", runtime.NumGoroutine())
	wg.Wait()
}

// Total CPU: 4
// Goroutines: 1
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine // Worker 2 goroutine
// Worker 2 goroutine // Worker 2 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine // Worker 1 goroutine
// Goroutines: 3
