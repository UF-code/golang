package main

import "fmt"

func sender(channel chan<- int) {
	channel <- 42
}

func receiver(channel <-chan int) {
	fmt.Println(<-channel)
}

func main() {

	channel := make(chan int)

	go sender(channel)

	go receiver(channel)
	// 42

	fmt.Println("Done")
	// Done
}
