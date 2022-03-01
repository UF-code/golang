package main

import "fmt"

func main() {
	sayHi("John Doe")        // Hi, John Doe
	sayHi_2("Frederick", 42) // Hi, Frederick Outside is 42 degree.
}

// single argument
func sayHi(name string) {
	fmt.Println("Hi,", name)
}

// multiple argument
func sayHi_2(name string, degree int) {
	fmt.Println("Hi,", name, "Outside is", degree, "degree.")
}
