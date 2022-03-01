package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func worker_1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Worker 1: ", i)
	}
	wg.Done()
}

func worker_2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Worker 2: ", i)
	}
	wg.Done()
}

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	// OS              windows
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	// ARCH            amd64
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	// CPUs            4
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// Goroutines      1

	wg.Add(1)
	go worker_1()
	wg.Add(1)
	go worker_2()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	// CPUs             4
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// Goroutines       3
	wg.Wait()
}
