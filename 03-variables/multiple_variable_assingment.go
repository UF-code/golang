package main

import "fmt"

func main() {
	var (
		multiple_var_integer  int = 10
		multiple_var_integer2     = 20

		multiple_var_string  string = "Hello"
		multiple_var_string2        = "World"

		multiple_var_bool  bool = true
		multiple_var_bool2      = false
	)

	fmt.Println(multiple_var_integer, multiple_var_integer2) // 10 20
	fmt.Println(multiple_var_string, multiple_var_string2)   // Hello World
	fmt.Println(multiple_var_bool, multiple_var_bool2)       // true false
}
