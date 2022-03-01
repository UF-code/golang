package main

import "fmt"

func main() {
	const (
		multiple_const_integer  int = 10
		multiple_const_integer2     = 20

		multiple_const_string  string = "Hello"
		multiple_const_string2        = "World"

		multiple_const_bool  bool = true
		multiple_const_bool2      = false
	)

	fmt.Println(multiple_const_integer, multiple_const_integer2) // 10 20
	fmt.Println(multiple_const_string, multiple_const_string2)   // Hello World
	fmt.Println(multiple_const_bool, multiple_const_bool2)       // true false
}
